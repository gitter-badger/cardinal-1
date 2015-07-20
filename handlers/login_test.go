package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	viper "github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

func initTestDB() *mgo.Database {
	var db *mgo.Database

	viper.AddConfigPath("../")
	viper.SetConfigFile("config.toml")
	viper.ReadInConfig()

	dialURL := "mongodb://" + viper.GetString("database.user") + ":" + viper.GetString("database.password") + "@" + viper.GetString("database.ip") + ":" + viper.GetString("database.port") + "/" + viper.GetString("database.name")
	session, err := mgo.Dial(dialURL)
	if err != nil {
		fmt.Println(err.Error())
	}

	db = session.DB(viper.GetString("dbname"))
	return db
}

func TestSignupHandler(t *testing.T) {
	db := initTestDB()

	mU, mE := json.Marshal(User{Username: "test", Password: []byte("test")})
	if mE != nil {
		t.Fatal("Marshal Error")
	}

	req, _ := http.NewRequest("POST", "/user/signup", bytes.NewReader(mU))
	res := httptest.NewRecorder()

	SignupHandler(res, req, db.C("users"))

	if res.Code != http.StatusForbidden {
		t.Fatalf("Expected: %v Got: %v", http.StatusForbidden, res.Code)
	}
}

func TestLoginHandler(t *testing.T) {
	db := initTestDB()

	mU, mE := json.Marshal(User{Username: "test", Password: []byte("test")})
	if mE != nil {
		t.Fatal("Marshal Error")
	}

	req, _ := http.NewRequest("GET", "/user/login/", bytes.NewReader(mU))
	res := httptest.NewRecorder()

	LoginHandler(res, req, db.C("users"))

	t.Log(res.Body)

	if res.Code != 302 {
		t.Fatalf("Expected: %v: Got: %v", "200", res.Code)
	}

}
