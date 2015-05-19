package main

import (
    "fmt"
    "net/http"
    "time"
    "os"
    "log"
    "github.com/gorilla/mux"
)

var logger log.Logger

func errCheck(err error) {
    if err != nil {
        logger <- err.Error()
    }
}

func getLogFile() *os.File {
    var logFile *os.File
    var fileErr error
    filename := "logs/" + time.Now().Format("01-02-2006") + "-http.log"

    if _, err := os.Stat(filename); err == nil {
        fmt.Println("File exists")
        logFile, fileErr = os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0660)
        logFile.WriteString("\n")
        if fileErr  != nil {
            fmt.Println(fileErr.Error())
        }
    } else {
        _, pErr := os.Stat("logs/")

        if os.IsNotExist(pErr) {
            os.Mkdir("logs/", 0777)
        }

        logFile, fileErr = os.Create(filename)
        if fileErr  != nil {
            fmt.Println(fileErr.Error())
        }
    }

    return logFile
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "../client/index.html")
}

func main() {
    fmt.Println("Server starting")
    logFile = getLogFile()
    mwriter := io.MultiWriter(os.Stdout, logFile)

    logger = log.New(mwriter, " [CARD-MANAGER] ", log.Ldate|log.Ltime)
    router := mux.NewRouter()

    router.HandleFunc("/login", loginHandler).Methods("POST")
    router.HandleFunc("/signup", signupHandler).Methods("POST")
    router.HandleFunc("/")

    router.PathPrefix("/").Handler(http.FileServer(http.Dir("../client/")))

    logger <- "Server ready"
    m.Run()
}
