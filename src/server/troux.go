package main

import (
        "fmt"
        "encoding/json"
        "net/http"
        "io/ioutil"
        "github.com/gorilla/mux"
        "time"
        "os"
        "bytes"
        "database/sql"
        _ "github.com/denisenkom/go-mssqldb"
        "strings"
        // _ "code.google.com/p/odbc"
)


var types = []string{"application", "server", "hardware", "softwareproduct", "softwareversion", "hardwareproduct", "hardwaremodel", "organization", "database", "server", "softwaremodule"}

type ResFromTroux struct {
    Model []json.RawMessage
}

type ResFromApp struct {
    Message string
    TrouxData json.RawMessage
}

type SqlResponse struct {
    Columns []string
    Rows [][]string
}

func log(message chan string) {
    var f *os.File;
    var fileError error;

    f, fileError = os.Create("logs/troux_" + time.Now().Format("2006-01-02") + ".log")
    if fileError != nil {
        fmt.Println(fileError)
    }

    if f == nil {
        f, fileError = os.Create("logs/troux_" + time.Now().Format("2006-01-02") + ".log")
        if fileError != nil {
            fmt.Println(fileError)
        }
    }

    defer f.Close()

    for {
        msg := <-message
        f.WriteString(time.Now().Format(time.ANSIC) + ": " + msg + "\n")
    }
}

func check(err error, message chan string) {
   if err != nil {
       msg := "Error: " + err.Error()
       message <- msg
   }
}

func sqlHandler(w http.ResponseWriter, r *http.Request, logger chan string, db *sql.DB) {
    vars := mux.Vars(r)
    euid := vars["euid"]
    query := vars["query"]
    statement := query + " FROM tcx3"
    logger <- euid + " Ran query: " + statement
    var sqlResp SqlResponse
    var colErr error

    rows, queryErr := db.Query(statement)
    if queryErr != nil {
        fmt.Println("Query Error: " + queryErr.Error())
        check(queryErr, logger)
    }
    defer rows.Close()

    sqlResp.Columns, colErr = rows.Columns()
    if colErr != nil {
        fmt.Println("ColErr = " + colErr.Error())
        check(colErr, logger)
    }

    // Result is your slice string.
    rawResult := make([][]byte, len(sqlResp.Columns))
    var result []string

    dest := make([]interface{}, len(sqlResp.Columns)) // A temporary interface{} slice
    for i, _ := range rawResult {
        dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
    }

    leng := len(sqlResp.Columns)
    count := 0

    for rows.Next() {
        err := rows.Scan(dest...)
        if err != nil {
            check(err, logger)
            return
        }


        for _, raw := range rawResult {
            count++
            if raw == nil {
                result = append(result, "\\N")
            } else {
                result = append(result, string(raw))
            }

            if count == leng {
                count = 0
                sqlResp.Rows = append(sqlResp.Rows, result)
                result = nil
            }
        }
    }

    numberOfResults := len(sqlResp.Rows) * len(sqlResp.Rows[0])
    logger <- "And they recieved: " + string(numberOfResults) + " no. of results"

    json.NewEncoder(w).Encode(sqlResp)
}

func searchThroughTrouxHandler( w http.ResponseWriter, r *http.Request, logger chan string, searchChannel chan *http.Response) {
    vars := mux.Vars(r)
    name := vars["searchTerm"]
    fmt.Println("Searching for " + name)
    urlName := strings.Replace(name, " ", "%20", -1)
    var responses ResFromTroux
    for i := 0; i < len(types); i++ {
        go func(name string, i int){
            fmt.Println("Searching for " + name + " in " + types[i])
            resp, err := http.Get("http://troux:troux@10.254.205.7:8088/tip/rest/v1/model/" + types[i] + "?select=name&filter=name%20.contains%20" + urlName)
            check(err, logger)

            searchChannel <- resp
        }(name, i)
    }

    for i := 0; i < len(types); i++ {
            resp := <- searchChannel

            respBody, _ := ioutil.ReadAll(resp.Body)
            var jsonHolder ResFromTroux
            UnmarshalError := json.Unmarshal(respBody, &jsonHolder)
            check(UnmarshalError, logger)

            if len(jsonHolder.Model) > 0 {
                for i := range jsonHolder.Model {
                    responses.Model = append(responses.Model, jsonHolder.Model[i])
                }
            }
    }


    json.NewEncoder(w).Encode(responses)
}

func getFullAppFromTrouxHandler( w http.ResponseWriter, r *http.Request,  logger chan string) {
    vars := mux.Vars(r)
    id := vars["ID"]
    typ := vars["TYPE"]
    url := "http://troux:troux@10.254.205.7:8088/tip/rest/v1/model/" + typ + "/" + id + "?select=" +
    "name," +
    "description," +
    "ITOwner.name," +
    "businessowner.name,businessowner.officephonenumber,businessowner.emailaddress," +
    "owningOrganization.name," +
    "category.name," +
    "logindate," +
    "providedBusinessFunction.name," +
    "numberOfUsers," +
    "productionDate," +
    "retirementDate," +
    "requiredSkills.name," +
    "requiredInterfaces.name,providedInterfaces.name," +
    "recommendationDate,recommendationComments," +
    "recurringCost,recurringCostInterval,initialDeploymentCost," +
    "annualSoftwareCosts,annualHardwareCosts,annualInternalStaffCosts,annualExternalStaffCosts,annualOthercosts," +
    "sessions"

    resp, err := http.Get(url)
    check(err, logger)

    logger <- "Searched for id: " + id + " " + resp.Status

    respBody := json.NewDecoder(resp.Body)
    var jsonHolder ResFromTroux
    jsonErr := respBody.Decode(&jsonHolder)
    check(jsonErr, logger)

    json.NewEncoder(w).Encode(jsonHolder)
}

func getSpecificFieldFromTrouxHandler( w http.ResponseWriter, r *http.Request,  logger chan string) {
    vars := mux.Vars(r)
    id := vars["ID"]
    typ := vars["TYPE"]
    reqFields := vars["REQFIELDS"]
    fmt.Println("Req Fields Found")
    url := "http://troux:troux@10.254.205.7:8088/tip/rest/v1/model/" + typ + "/" + id + "?select=" + reqFields

    resp, err := http.Get(url)
    check(err, logger)

    respBody := json.NewDecoder(resp.Body)
    var jsonHolder ResFromTroux
    jsonErr := respBody.Decode(&jsonHolder)
    check(jsonErr, logger)

    json.NewEncoder(w).Encode(jsonHolder)
}

func postToTrouxHandler( w http.ResponseWriter, r *http.Request, logger chan string) {
    var appData ResFromApp
    vars := mux.Vars(r)
    id := vars["ID"]
    typ := vars["TYPE"]

    respBody := json.NewDecoder(r.Body)

    err := respBody.Decode(&appData)
    check(err, logger)

    logger <- appData.Message

    url := "http://troux:troux@10.254.205.7:8088/tip/" + typ + "/" + id

    req, reqErr := http.NewRequest("POST", url, bytes.NewBuffer([]byte(appData.TrouxData)))
    check(reqErr, logger)

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, sendErr := client.Do(req)
    check(sendErr, logger)

    defer resp.Body.Close()

    body, trouxError := ioutil.ReadAll(resp.Body)
    check(trouxError, logger)

    logger <- "POST MESSAGE " + string(body)
}

func loadTestHandler(w http.ResponseWriter, r *http.Request, logger chan string) {
    http.ServeFile(w, r, "views/index.html")
}

func main() {
    fmt.Println("Troux/SQL Interface Starting")

    logger := make(chan string)
    go log(logger)
    defer close(logger)

    searchChannel := make(chan *http.Response)
    defer close(searchChannel)

    logger <- "Channels Successfully Created"
    var db *sql.DB
    var dbErr error;
    connString := "server=n060sqlt04.kroger.com;user id=svcConfluence;password=7qD8PyDh;port=1675;database=ISS_Confluence_Test"
    db, dbErr = sql.Open("mssql", connString)
    check(dbErr, logger)
    defer db.Close()
    if db == nil {
        logger <- "DB Connection Failed"
    } else {
        logger <- "DB Connection Successful"
    }

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/troux_search/{searchTerm}", func(w http.ResponseWriter, r *http.Request){
        searchThroughTrouxHandler(w, r, logger, searchChannel)
        })
    logger <- "Search Handler Registered"
    router.HandleFunc("/troux_get/{TYPE}/{ID}/{REQFIELDS}", func(w http.ResponseWriter, r *http.Request){
        getSpecificFieldFromTrouxHandler(w, r, logger)
        })
    router.HandleFunc("/troux_get/{TYPE}/{ID}", func(w http.ResponseWriter, r *http.Request){
        getFullAppFromTrouxHandler(w, r, logger)
        })
    logger <- "Get Handlers Registered"
    router.HandleFunc("/troux_post/{TYPE}/{ID}", func(w http.ResponseWriter, r *http.Request){
        postToTrouxHandler(w, r, logger)
        })
    logger <- "Post Handler Registered"
    router.HandleFunc("/sql-query/{euid}/{query}", func(w http.ResponseWriter, r *http.Request){
        sqlHandler(w, r, logger, db)
        })
    logger <- "SQL Handler Registered"
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        loadTestHandler(w, r, logger)
        })

    router.Handle("/views/{file}", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))
    router.Handle("/css/{file}", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
    router.Handle("/js/{file}", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
    logger <- "Serving Static Files"

    logger <- "Ready for Requests"
    fmt.Println("Ready")
    http.ListenAndServe(":8585", router)
}
