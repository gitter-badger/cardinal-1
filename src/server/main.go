package main

import (
    "fmt"
    "net/http"
    "time"
    "os"
    "log"
    "io"
    "bufio"
    "github.com/gorilla/mux"
    "gopkg.in/mgo.v2"
    "github.com/suapapa/go_sass"
    "strings"
)

var logger *log.Logger
var db *mgo.Session
var magicCollection *mgo.Collection
var hearthStoneCollection *mgo.Collection
var userCollection *mgo.Collection

func loadDbProperties() (string, string){
    propFile, err := os.Open("db.properties")
    errCheck(err)
    scanner := bufio.NewScanner(propFile)
    var keys = make(map[string]string)

    for scanner.Scan() {
        line := scanner.Text()
        split := strings.Split(line, "=")
        keys[split[0]] = split[1]
    }

    url :=  "mongodb://" + keys["user"] + ":" + keys["password"] + "@" + keys["ip"] + ":" + keys["port"] + "/" + keys["dbName"]
    logger.Println("Connecting to DB: " + url)
    return url, keys["dbName"]
}

func initDB() {
    var err error
    dialURL, dbname := loadDbProperties()
    db, err = mgo.Dial(dialURL)

    errCheck(err)
    if err != nil {
        logger.Panic(err)
    }

    magicCollection = db.DB(dbname).C("magiccards")
    hearthStoneCollection = db.DB(dbname).C("hearthstonecards")
    userCollection = db.DB(dbname).C("users")
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
        if fileErr  != nil {
            fmt.Println(fileErr.Error())
        }
    } else {
        _, pErr := os.Stat("../logs/")

        if os.IsNotExist(pErr) {
            os.Mkdir("../logs/", 0777)
        }

        logFile, fileErr = os.Create(filename)
        if fileErr  != nil {
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

    var sc sass.Compiler
    sc.CompileFolder("../client/scss", "../client/css")
    logger.Println("Sass compiled")
    initDB()
    defer db.Close()
    logger.Println("DB Initialized")

    router := mux.NewRouter()

    router.HandleFunc("/login", loginHandler).Methods("POST")
    router.HandleFunc("/signup", signupHandler).Methods("POST")
    router.HandleFunc("/", indexHandler)

    router.PathPrefix("/").Handler(http.FileServer(http.Dir("../client/")))

    logger.Println("Router Created")
    logger.Println("Server ready")
    http.ListenAndServe(":3001", router)
}
