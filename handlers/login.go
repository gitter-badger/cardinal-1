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
	hasher     = md5.New()
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
	Username    string             `json:"username"`
	Password    []byte             `json:"-"`
	Token       string             `json:"authToken"`
	DashItems   []DashItem         `json:"dashitems"`
	Collections []cards.Collection `json:"collections"`
}

func generateToken(un string) {
	unique := false

	for !unique {
		buffer := make([]byte, keySize)
		_, err := rand.Read(buffer)
		if err != nil {
			logger.Fatalln(err.Error())
		}

		token := base64.URLEncoding.EncodeToString(buffer)
		_, exists := ssoExpires[token]

		if !exists {
			unique = true
		}
	}
}

// This is meant to be run as a go-routine on a timer.
func cleanOutSso() {
	for user, token := range sso {
		if time.Since(ssoExpires[token]).Minutes() > 30 {
			delete(ssoExpires, token)
			delete(sso, user)
		}
	}
}

func isLoggedInSSO(un string) bool {
	if sso[un] != "" {
		return true
	}
	return false
}

// LoginHandler will log a user in if they exist otherwise will return an error.
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
		marshaledU, merr := json.Marshal(udb)
		if merr != nil {
			logger.Error(merr.Error())
		}

		logger.Info("User " + u.Username + " has successfully logged in.")
		w.Write(marshaledU)
	} else {
		logger.Info("User " + u.Username + " failed login attempt.")
		w.WriteHeader(http.StatusForbidden)
	}
}

// SignupHandler accepts a json formatted user and will update them into the given collection.
func SignupHandler(w http.ResponseWriter, r *http.Request, collection *mgo.Collection) {
	decoder := json.NewDecoder(r.Body)
	var u User
	logger.Debug("signup handler called")

	err := decoder.Decode(&u)
	if err != nil {
		logger.Error(err.Error())
	}

	u.Password = hasher.Sum(u.Password)
	u.DashItems = append(u.DashItems, DashItem{Img: "/img/default.jpg", Title: "Dash Item Title", Content: "This is your default dash item! You can create your own by choosing \"Edit Dash\" from the side Menu!"})
	derr := collection.Insert(u)
	if derr != nil {
		w.WriteHeader(http.StatusForbidden)
	} else {
		marshaledU, merr := json.Marshal(u)
		if merr != nil {
			logger.Error(merr.Error())
		}

		logger.Info("User " + u.Username + " has successfully signed up.")
		w.Write(marshaledU)
	}
}
