package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "net/http"
    "io/ioutil"
)

func main() {
    /*fileServer := http.FileServer(http.Dir("../web/html/"))
    http.Handle("/", fileServer)*/ 
    _ = initConnection()
    http.HandleFunc("/newData", newDataHandler)

    // Take this out in prod!
    http.HandleFunc("/testRequest", testRequest)

    fmt.Printf("Starting Fishtank Monitor Server\n")
    if err:= http.ListenAndServe(":8081",nil); err != nil {
        log.Fatal(err)
    }
}

// Post method for js
func newDataHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/newData" {
        http.Error(w, "400 Bad Request.", http.StatusBadRequest)
        return
    }

    if r.Method != "POST" {
        http.Error(w, "Method is not supported.", http.StatusBadRequest)
        return
    }

    // Ensure user is authenticated to be able to post data
    // CREATE KEY IN /app/key 
    if isRequestFromAuth(r) {
        body, _ := ioutil.ReadAll(r.Body)
        fmt.Println("Received datapoint: ", string(body))
        // Parse body into database
        sendToDatabase(body)
    } else {
        fmt.Println("Denied request; no key matches key file.")
    }
}

func isRequestFromAuth(r *http.Request) bool {
    if r.URL.Query()["key"] != nil {
        authKey := r.URL.Query()["key"]
        dat, err := os.ReadFile("/app/key")
        if err == nil {
            if strings.TrimSpace(string(dat)) == authKey[0] {
                return true
            }
        }
    }
    return false
}

func testRequest(w http.ResponseWriter, r *http.Request) {
    // Test whatever I need to while writing code
    fmt.Fprintf(w, r.URL.Query()["test"][0])
}

func sendDatabaseData(json []byte){
    // Parse json into type structure
}
