package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/suapapa/go_sass"
	"gopkg.in/mgo.v2"
)

var (
	logger                *log.Logger
	db                    *mgo.Session
	magicCollection       *mgo.Collection
	hearthStoneCollection *mgo.Collection
	userCollection        *mgo.Collection
)

func loadDbProperties() (string, string) {
	propFile, err := os.Open("db.properties")
	if err != nil {
		logger.Fatalln(err.Error())
	}
	scanner := bufio.NewScanner(propFile)
	var keys = make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "=")
		keys[split[0]] = split[1]
	}

	url := "mongodb://" + keys["user"] + ":" + keys["password"] + "@" + keys["ip"] + ":" + keys["port"] + "/" + keys["dbName"]
	logger.Println("Connecting to DB: " + url)
	return url, keys["dbName"]
}

func initDB() {
	var err error
	dialURL, dbname := loadDbProperties()
	db, err = mgo.Dial(dialURL)

	errCheck(err)
	if err != nil {
		logger.Fatalln(err.Error())
	}

	hearthStoneCollection = db.DB(dbname).C("hearthstonecards")
	userCollection = db.DB(dbname).C("users")
	magicCollection = db.DB(dbname).C("magiccards")
}

func errCheck(err error) {
	if err != nil {
		logger.Println(err.Error())
	}
}

func getLogFile() *os.File {
	var logFile *os.File
	var fileErr error
	filename := "../logs/" + time.Now().Format("01-02-2006") + "-http.log"

	if _, err := os.Stat(filename); err == nil {
		fmt.Println("File exists")
		logFile, fileErr = os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0660)
		logFile.WriteString("\n")
		if fileErr != nil {
			fmt.Println(fileErr.Error())
		}
	} else {
		_, pErr := os.Stat("../logs/")

		if os.IsNotExist(pErr) {
			os.Mkdir("../logs/", 0777)
		}

		logFile, fileErr = os.Create(filename)
		if fileErr != nil {
			fmt.Println(fileErr.Error())
		}
	}

	return logFile
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/views/index.html")
}

func main() {
	fmt.Println("Server starting")
	logFile := getLogFile()
	mwriter := io.MultiWriter(os.Stdout, logFile)

	logger = log.New(mwriter, "[CARD-MANAGER] ", log.Ldate|log.Ltime)
	logger.Println("Logger Initialized")

	logger.Println("Compiling Sass")
	var sc sass.Compiler
	sc.CompileFolder("../client/sass", "../client/css")
	logger.Println("Sass compiled")
	logger.Println("Connecting to DB")
	initDB()
	defer db.Close()
	logger.Println("DB Initialized")

	logger.Println("Creating Router")
	router := mux.NewRouter()

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/signup", signupHandler).Methods("POST")
	router.HandleFunc("/search/{game}/{name}", searchHandler).Methods("GET")
	router.HandleFunc("/", indexHandler)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../client/")))

	logger.Println("Router Created")
	logger.Println("Server ready")
	http.ListenAndServe(":3001", router)
}
