package main

import (
    "fmt"
    "log"
    "os"
    "database/sql"
    "strings"
    "net/http"
    "io/ioutil"
)

func main() {
    db := initConnection()
    http.HandleFunc("/newData", newDataWrapper(db))

    // Take this out in prod!
    http.HandleFunc("/testRequest", testRequest)

    fmt.Printf("Starting Fishtank Monitor Server\n")
    if err:= http.ListenAndServe(":8081",nil); err != nil {
        log.Fatal(err)
    }
}


func newDataWrapper(db *sql.DB) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request) {
        body := newDataHandler(w,r)
        // Parse body into database
        sendToDatabase(db, body)
    }

}
// Post method for js
func newDataHandler(w http.ResponseWriter, r *http.Request) []byte {
    if r.URL.Path != "/newData" {
        http.Error(w, "400 Bad Request.", http.StatusBadRequest)
        return nil
    }

    if r.Method != "POST" {
        http.Error(w, "Method is not supported.", http.StatusBadRequest)
        return nil
    }

    // Ensure user is authenticated to be able to post data
    // CREATE KEY IN /app/key 
    if isRequestFromAuth(r) {
        body, _ := ioutil.ReadAll(r.Body)
        fmt.Println("Received datapoint: ", string(body))
        return body
    } else {
        fmt.Println("Denied request; no key matches key file.")
        return nil
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

