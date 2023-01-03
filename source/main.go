package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
)


// TODO: SQL server interaction
type TankInfo struct {
    Tank_ID int
    Temperature float32
    Ph float32
}

func main() {
    /*fileServer := http.FileServer(http.Dir("../web/html/"))
    http.Handle("/", fileServer)*/

    http.HandleFunc("/newData", newDataHandler)
    http.HandleFunc("/hello", helloHandler)

    fmt.Printf("Starting Fishtank Monitor Server\n")
    if err:= http.ListenAndServe(":8080",nil); err != nil {
        log.Fatal(err)
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }
    fmt.Fprintf(w,"Hello!")
}

func newDataHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/newData" {
        http.Error(w, "400 Bad Request.", http.StatusBadRequest)
        return
    }

    if r.Method != "POST" {
        http.Error(w, "Method is not supported.", http.StatusBadRequest)
        return
    }
    body, _ := ioutil.ReadAll(r.Body)
    fmt.Println(string(body))
}

// Send to database

func sendDatabaseData(json []byte){
    // Parse json into type structure

}
