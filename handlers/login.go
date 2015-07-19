package handlers

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ChasingLogic/cardinal/cards"
	logger "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	hasher = md5.New()
	// Storing sso info in memory to increase speed and security. Need to find a way to prevent "hijacking"
	ssoExpires = make(map[string]time.Time)
	sso        = make(map[string]string)
	keySize    = 32
)

// DashItem is the basic form of a "Dash Card" which displays various information to the user.
type DashItem struct {
	Img     string `json:"img"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// User is to hold the user information in our DB
type User struct {
	Username string `json:"username"`
	Password []byte `json:"-"`
	// The front-end will need to store this as a cookie
	Token       string             `json:"token"`
	DashItems   []DashItem         `json:"dashitems"`
	Collections []cards.Collection `json:"collections"`
}

func generateToken(un string) (string, error) {
	var token string
	logger.Debug("Generating token for " + un)

	unique := false
	for !unique {
		buffer := make([]byte, keySize)
		_, err := rand.Read(buffer)
		if err != nil {
			return "", err
		}

		token = base64.URLEncoding.EncodeToString(buffer)
		logger.Debug("Generated token: " + token)
		_, exists := ssoExpires[token]

		if !exists {
			logger.Debug("Token was unique")
			unique = true
		}

		logger.Debug("Token was not unique")
	}

	sso[token] = un
	ssoExpires[token] = time.Now()

	return token, nil
}

// This is meant to be run as a go-routine on a timer. Since we are storing sso tokens in memory we need to make sure
// we aren't leaking memory like the Titanic
func cleanOutSso() int {
	var i = 0
	for token := range sso {
		if time.Since(ssoExpires[token]).Minutes() > 30 {
			delete(ssoExpires, token)
			delete(sso, token)
			i++
		}
	}

	return i
}

func tokenNotExpired(token string) bool {
	timeSinceTouched := ssoExpires[token]
	if time.Since(timeSinceTouched).Minutes() > 30 {
		return false
	}

	// 'touch' the token so it's session is renewed since we have a successful login.
	ssoExpires[token] = time.Now()
	return true
}

// LoginHandler will log a user in if they exist otherwise will return an error.
// We just send the full user as a json object instead of splitting it across a url param and json object as is the norm.
func LoginHandler(w http.ResponseWriter, r *http.Request, collection *mgo.Collection) {
	decoder := json.NewDecoder(r.Body)
	var u User
	var udb User
	logger.Debug("login handler called")

	err := decoder.Decode(&u)
	if err != nil {
		logger.Error(err.Error())
	}

	derr := collection.Find(bson.M{"username": u.Username}).One(&udb)
	if derr != nil {
		logger.Error(derr.Error())
	}

	u.Password = hasher.Sum(u.Password)
	if string(u.Password) == string(udb.Password) {
		token, err := generateToken(u.Username)
		if err != nil {
			logger.Error(err.Error())
		}

		udb.Token = token

		marshaledU, merr := json.Marshal(udb)
		if merr != nil {
			logger.Error(merr.Error())
		}

		logger.Info("User " + u.Username + " has successfully logged in.")
		logger.Debug("Token generated for " + u.Username + " is " + token)
		w.Write(marshaledU)
	} else {
		logger.Info("User " + u.Username + " failed login attempt.")
		w.WriteHeader(http.StatusForbidden)
	}
}

// SignupHandler accepts a json formatted user then will add the defaults to the user and will update them into the given collection.
// It then returns this new user to the front-end
func SignupHandler(w http.ResponseWriter, r *http.Request, collection *mgo.Collection) {
	decoder := json.NewDecoder(r.Body)
	var u User
	logger.Debug("signup handler called")

	err := decoder.Decode(&u)
	if err != nil {
		logger.Error(err.Error())
	}

	u.Password = hasher.Sum(u.Password)
	u.DashItems = append(u.DashItems, DashItem{Img: "/img/defaultDashItem.jpg", Title: "Dash Item Title", Content: "This is your default dash item! You can create your own by choosing \"Edit Dash\" from the side Menu!"})
	derr := collection.Insert(u)
	if derr != nil {
		w.WriteHeader(http.StatusForbidden)
	}

	marshaledU, merr := json.Marshal(u)
	if merr != nil {
		logger.Error(merr.Error())
	}

	logger.Info("User " + u.Username + " has successfully signed up.")
	w.Write(marshaledU)
}

// SSOHandler will handle SSO Requests and updates
// SSOHandler will live at /user/sso and take a parameter of ?token=
func SSOHandler(w http.ResponseWriter, r *http.Request, collection *mgo.Collection) {
	token := r.FormValue("token")
	un, exists := sso[token]
	if exists && tokenNotExpired(token) {
		var u User
		collection.Find(bson.M{"username": un}).One(&u)
		marshaledU, merr := json.Marshal(u)
		if merr != nil {
			logger.Error(merr.Error())
		}

		w.Write(marshaledU)
	}
}
