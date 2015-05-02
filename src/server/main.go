package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "time"
    "os"
    "encoding/json"
    // "io/ioutil"
)

type User struct {
    Username string
    Password string
}


var logger = make(chan string)

func errCheck(err error) {
    if err != nil {
        logger <- err.Error()
    }
}

func log(logger chan string) {
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

    defer logFile.Close()

    for {
        message := <- logger
        logFile.WriteString(time.Now().Format("Jan _2 15:04:05") + " " + message)
    }
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var u User
    err := decoder.Decode(&u)
    errCheck(err)
    // Do Query stuff here
    // This is fake auth
    if u.Username == "ChasingLogic" && u.Password == "test" {
        w.Write([]byte("TRUE"))
    } else {
        w.Write([]byte("FALSE"))
    }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "../client/index.html")
}

func main() {
    fmt.Println("Server starting")

    go log(logger)

    router := mux.NewRouter()
    router.HandleFunc("/", indexHandler)
    router.HandleFunc("/login", loginHandler)

    router.PathPrefix("/").Handler(http.FileServer(http.Dir("../client/")))

    logger <- "Server ready"

    http.ListenAndServe(":8080", router)
}
