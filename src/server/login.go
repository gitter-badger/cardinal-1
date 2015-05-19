package main

import (
    "net/http"
    "encoding/json"
    )

func loginHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var u User
    err := decoder.Decode(&u)
    errCheck(err)
    // Do Query stuff here
    // This is fake auth
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
    
}
