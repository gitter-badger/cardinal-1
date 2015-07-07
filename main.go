package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/ChasingLogic/configoslurper"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

var (
	settings              map[string]string
	logger                = logrus.New()
	slurper               = configoslurper.GetBasicSlurper("application.properties")
	db                    *mgo.Session
	magicCollection       *mgo.Collection
	hearthStoneCollection *mgo.Collection
	userCollection        *mgo.Collection
)

func initDB() {
	var err error

	dialURL := "mongodb://" + settings["dbuser"] + ":" + settings["dbpassword"] + "@" + settings["dbaddress"] + ":" + settings["dbport"] + "/" + settings["dbname"]
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

func loggerInit() (*os.File, logrus.Level) {
	var logFile *os.File
	var fileErr error
	filename := "cardcollector.out"

	fmt.Println("Log File: " + filename)

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

	switch strings.ToLower(settings["loglevel"]) {

	case "debug":
		return logFile, logrus.DebugLevel

	case "info":
		return logFile, logrus.InfoLevel

	case "warn":
		return logFile, logrus.WarnLevel

	case "error":
		return logFile, logrus.ErrorLevel

	case "fatal":
		return logFile, logrus.FatalLevel

	case "panic":
		return logFile, logrus.PanicLevel

	default:
		return logFile, logrus.InfoLevel

	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Incoming request recieved")
	http.ServeFile(w, r, "client/views/index.html")
}

func main() {
	fmt.Println("Server starting")

	var serr error
	settings, serr = slurper.Slurp()
	if serr != nil {
		logger.Fatalln(serr.Error())
	}

	logFile, logLevel := loggerInit()
	mwriter := io.MultiWriter(os.Stdout, logFile)

	logger.Out = mwriter
	logger.Level = logLevel
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "01/02/2006 15:04:05",
	}
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

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("client/")))
	logger.Debug("serving static client files")

	logger.Debug("Router created")
	if settings["port"] == "" {
		settings["port"] = "8080"
	}
	logger.Info("Server ready on port " + settings["port"])
	http.ListenAndServe(":"+settings["port"], router)
}
