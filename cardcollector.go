package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/ChasingLogic/cardinal/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	logger "github.com/Sirupsen/logrus"
	viper "github.com/spf13/viper"
)

var (
	session        *mgo.Session
	db             *mgo.Database
	userCollection *mgo.Collection
)

func initDB() {
	var err error

	dialURL := "mongodb://" + viper.GetString("database.user") + ":" + viper.GetString("database.password") + "@" + viper.GetString("database.ip") + ":" + viper.GetString("database.port") + "/" + viper.GetString("database.name")
	logger.Debug("Connecting to mongodb @ " + dialURL)
	session, err = mgo.Dial(dialURL)

	if err != nil {
		logger.Fatalln(err.Error())
	}

	db = session.DB(viper.GetString("dbname"))
	userCollection = db.C(viper.GetString("database.userCollection"))
}

func loggerInit() (*os.File, logger.Level) {
	var logFile *os.File
	var fileErr error
	filename := "cardcollector.out"

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

	switch strings.ToLower(viper.GetString("server.loglevel")) {

	case "debug":
		return logFile, logger.DebugLevel

	case "info":
		return logFile, logger.InfoLevel

	case "warn":
		return logFile, logger.WarnLevel

	case "error":
		return logFile, logger.ErrorLevel

	case "fatal":
		return logFile, logger.FatalLevel

	case "panic":
		return logFile, logger.PanicLevel

	default:
		return logFile, logger.InfoLevel

	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debug("root handler called")
	http.ServeFile(w, r, "client/views/index.html")
}

func main() {

	viper.AddConfigPath("./")
	viper.SetConfigFile("config.toml")
	viper.ReadInConfig()

	logFile, logLevel := loggerInit()
	mwriter := io.MultiWriter(os.Stdout, logFile)

	logger.SetOutput(mwriter)
	logger.SetLevel(logLevel)
	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "01/02/2006 15:04:05",
	})

	logger.Info("Cardinal starting")
	logger.Info("Log file: " + logFile.Name())
	logger.Info("=========Settings===========")
	for k, v := range viper.AllSettings() {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			logger.Infof("%s", k)
			for sk, sv := range viper.GetStringMapString(k) {
				logger.Infof("\t%s = %s", sk, sv)
			}
		} else {
			logger.Infof("%s = %s", k, v)
		}

	}
	logger.Info("============================")

	logger.Info("connecting to mongodb")

	initDB()
	defer session.Close()

	logger.Info("mongodb connection successful")

	router := mux.NewRouter()

	router.HandleFunc("/userAuth/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, userCollection)
	}).Methods("POST")

	router.HandleFunc("/userAuth/signup", func(w http.ResponseWriter, r *http.Request) {
		handlers.SignupHandler(w, r, userCollection)
	}).Methods("POST")

	router.HandleFunc("/api/v1/cardSearch", func(w http.ResponseWriter, r *http.Request) {
		handlers.CardSearch(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/", indexHandler)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("client/")))

	logger.Info("Server ready on port " + viper.GetString("server.port"))
	http.ListenAndServe(viper.GetString("server.ip")+":"+viper.GetString("server.port"), router)
}
