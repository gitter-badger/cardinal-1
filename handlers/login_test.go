package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	viper "github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

func initTestDB() *mgo.Database {
	var db *mgo.Database

	configByte, ferr := ioutil.ReadFile("../config.toml")
	if ferr != nil {
		fmt.Println(ferr.Error())
	}

	viper.SetConfigType("toml")

	verr := viper.ReadConfig(bytes.NewReader(configByte))
	if verr != nil {
		fmt.Println(verr.Error())
	}

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

	if res.Code != 403 {
		t.Fatalf("Expected: %v Got: %v", 403, res.Code)
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

	if res.Code != 200 {
		t.Fatalf("Expected: %v: Got: %v", "200", res.Code)
	}
}

func TestSSOHandler(t *testing.T) {
	db := initTestDB()

	mU, mE := json.Marshal(User{Username: "test", Password: []byte("test")})
	if mE != nil {
		t.Fatal("Marshal Error")
	}

	req, _ := http.NewRequest("GET", "/user/login/", bytes.NewReader(mU))
	res := httptest.NewRecorder()

	LoginHandler(res, req, db.C("users"))

	if res.Code != 200 {
		t.Fatalf("Expected: %v: Got: %v", "200", res.Code)
	}

	var umU User
	umE := json.Unmarshal(res.Body.Bytes(), &umU)
	if umE != nil {
		t.Fatalf("Unable to unmarshal response back into a user.")
	}

	ssoReq, _ := http.NewRequest("POST", "/user/sso?token="+umU.Token, nil)
	ssoRes := httptest.NewRecorder()

	SSOHandler(ssoRes, ssoReq, db.C("users"))

	if ssoRes.Code != 200 {
		t.Fatalf("Unexpected response, Expected: 200 Got: %v\nResponse: %v", ssoRes.Code, ssoRes)
	}
}

func GetToken(db *mgo.Database) string {
	mU, mE := json.Marshal(User{Username: "test", Password: []byte("test")})
	if mE != nil {
		fmt.Println("Marshal Error")
	}

	req, _ := http.NewRequest("GET", "/user/login/", bytes.NewReader(mU))
	res := httptest.NewRecorder()

	LoginHandler(res, req, db.C("users"))

	if res.Code != 200 {
		fmt.Printf("Login failed. Expected: %v: Got: %v", "200", res.Code)
	}

	var umU User
	umE := json.Unmarshal(res.Body.Bytes(), &umU)
	if umE != nil {
		fmt.Println("Unable to unmarshal response back into a user.")
	}

	return umU.Token
}

func TestGenerateToken(t *testing.T) {
	db := initTestDB()

	firstToken := GetToken(db)
	secondToken := GetToken(db)

	t.Log("Tokens generated")

	if firstToken != secondToken {
		t.Fatalf("Expected tokens to match instead firstToken: %s, secondToken %s", firstToken, secondToken)
	}
}
