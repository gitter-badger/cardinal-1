package main

import (
	"crypto/md5"
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

var (
	hasher = md5.New()
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u User
	var udb User
	logger.Debug("login handler called")

	err := decoder.Decode(&u)
	errCheck(err)

	derr := userCollection.Find(bson.M{"username": u.Username}).One(&udb)
	errCheck(derr)

	u.Password = hasher.Sum(u.Password)
	if string(u.Password) == string(udb.Password) {
		marshaledU, merr := json.Marshal(udb)
		errCheck(merr)
		logger.Info("User " + u.Username + " has successfully logged in.")
		w.Write(marshaledU)
	} else {
		logger.Info("User " + u.Username + " failed login attempt.")
		w.WriteHeader(http.StatusForbidden)
	}
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u User
	logger.Debug("signup handler called")

	err := decoder.Decode(&u)
	errCheck(err)

	u.Password = hasher.Sum(u.Password)
	u.DashItems = append(u.DashItems, DashItem{Img: "/img/default.jpg", Title: "Dash Item Title", Content: "This is your default dash item! You can create your own by choosing \"Edit Dash\" from the side Menu!"})
	derr := userCollection.Insert(u)
	if derr != nil {
		errCheck(derr)
		w.WriteHeader(http.StatusForbidden)
	} else {
		marshaledU, merr := json.Marshal(u)
		errCheck(merr)
		logger.Info("User " + u.Username + " has successfully signed up.")
		w.Write(marshaledU)
	}
}
