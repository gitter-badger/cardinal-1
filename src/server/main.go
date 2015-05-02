package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "time"
    "os"
    // "io/ioutil"
)


var logger = make(chan string)

func errCheck(err error, logger chan string) {
    if err != nil {
        logger <- err.Error()
    }
}

func log(logger chan string) {
    var logFile *os.File
    var fileErr error
    filename := "logs/" + time.Now().Format("JANUARY_2_15_04_05") + "-http.log"

    if _, err := os.Stat(filename); err == nil {
        fmt.Println("File Found opening")
        logFile, fileErr = os.Open(filename)
        errCheck(fileErr, logger)
    } else {
        fmt.Println("FIle not found creating")
        logFile, fileErr = os.Create(filename)
        errCheck(fileErr, logger)
    }

    defer logFile.Close()

    for {
        message := <- logger
        logFile.WriteString(time.Now().Format("Jan _2 15:04:05") + " " + message)
    }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    logger <- "Index called"
    w.Write([]byte("Hello WOrld"))
}

func main() {
    fmt.Println("Server starting")
    fmt.Println("Channle made")

    go log(logger)
    fmt.Println("Logger started")

    fmt.Println("Hello World")

    router := mux.NewRouter()
    router.HandleFunc("/", indexHandler)
    fmt.Println("Router made")

    logger <- "Sometin  g"

    http.ListenAndServe(":8080", router)
}
