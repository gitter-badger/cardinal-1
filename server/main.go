package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ChasingLogic/configoslurper"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var (
	logger                = logrus.New()
	db                    *mgo.Session
	magicCollection       *mgo.Collection
	hearthStoneCollection *mgo.Collection
	userCollection        *mgo.Collection
)

func initDB() {
	var err error
	slurper := configoslurper.GetBasicSlurper("db.properties")
	settings := slurper.Slurp()
	logger.Debug(settings)

	dialURL := "mongodb://" + settings["user"] + ":" + settings["password"] + "@" + settings["ip"] + ":" + settings["port"] + "/" + settings["dbname"]
	logger.Info("Connecting to mongodb @ " + dialURL)
	db, err = mgo.Dial(dialURL)

	errCheck(err)
	if err != nil {
		logger.Fatalln(err.Error())
	}

	hearthStoneCollection = db.DB(settings["dbname"]).C("hearthstonecards")
	userCollection = db.DB(settings["dbname"]).C("users")
	magicCollection = db.DB(settings["dbname"]).C("magiccards")
}

func errCheck(err error) {
	if err != nil {
		logger.Warn(err.Error())
	}
}

func getLogFile() *os.File {
	var logFile *os.File
	var fileErr error
	filename := "../logs/" + time.Now().Format("01-02-2006") + "-http.log"

	if _, err := os.Stat(filename); err == nil {
		logFile, fileErr = os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0660)
		logFile.WriteString("\n")
		if fileErr != nil {
			fmt.Println(fileErr.Error())
			os.Exit(1)
		}

	} else {
		_, pErr := os.Stat("../logs/")

		if os.IsNotExist(pErr) {
			os.Mkdir("../logs/", 0777)
		}

		logFile, fileErr = os.Create(filename)
		if fileErr != nil {
			fmt.Println(fileErr.Error())
			os.Exit(1)
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

	logger.Out = mwriter
	logger.Level = logrus.DebugLevel
	logger.Formatter = new(logrus.TextFormatter)
	logger.Debug("logger initialized")

	logger.Info("connecting to mongodb")
	initDB()
	defer db.Close()
	logger.Info("mongodb connection successfull")

	logger.Debug("Creating router")
	router := mux.NewRouter()

	router.HandleFunc("/login", loginHandler).Methods("POST")
	logger.Debug("login handler registered")
	router.HandleFunc("/signup", signupHandler).Methods("POST")
	logger.Debug("signup handler registered")
	router.HandleFunc("/search/{game}/{name}", searchHandler).Methods("GET")
	logger.Debug("search handler registered")
	router.HandleFunc("/", indexHandler)
	logger.Debug("root handler registered")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../client/")))
	logger.Debug("serving static client files")

	logger.Debug("Router created")
	logger.Info("Server ready")
	http.ListenAndServe(":3001", router)
}
