package main

import (
    "net/http"
    "encoding/json"
    "fmt"
    "gopkg.in/mgo.v2/bson"
    "crypto/md5"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var u User
    var udb User
    hasher := md5.New()
    err := decoder.Decode(&u)
    errCheck(err)
    derr := userCollection.Find(bson.M{"username": u.Username}).One(&udb)
    fmt.Println("Doc Error")
    errCheck(derr)
    hashedUpPwd := hasher.Sum(u.Password)

    w.WriteHeader(http.StatusForbidden)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var u User
    err := decoder.Decode(&u)

}
