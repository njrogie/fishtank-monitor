package main

import (
    "fmt"
    "encoding/json"
    "github.com/cockroachdb/cockroach-go"
)

func initConnection() bool {
    fmt.Println("initializing connection.")
    
    return true
}

func sendToDatabase(jsonBytes []byte) {
    // Pack into known type
    var tInfo TankInfo
    err := json.Unmarshal(jsonBytes, tInfo)
    if err != nil {
        fmt.Println("Error parsing json from POST request.")
    }
    // Export the datapoint to the database

}
