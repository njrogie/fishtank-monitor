package main

import (
    "fmt"
    "time"
    //"strings"
    "encoding/json"
    "database/sql"
     _ "github.com/lib/pq"
    //"github.com/cockroachdb/cockroach-go/v2/crdb"
)

func initConnection() *sql.DB {
    fmt.Println("initializing connection.")
    // Test for now; later on improve connection complexity
    connStr := "host=localhost port=26257 dbname=mydb user=nick sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err == nil {
        // Do init stuff here
        fmt.Println("Connection made")
        return db
    } else {
        fmt.Println(fmt.Sprintf("Connection failed! Error:%s",err))
        return nil
    }
}

func sendToDatabase(db *sql.DB, jsonBytes []byte) {
    // Pack into known type
    var tInfo TankInfo
    err := json.Unmarshal(jsonBytes, &tInfo)
    if err != nil {
        fmt.Println("Error parsing json from POST request.")
        fmt.Println(err)
        return
    }
    datetime := time.Now()
    dt := datetime.Format("20060102 03:04:05 PM")
    fmt.Println(dt)
    _, err = db.Exec(fmt.Sprintf("INSERT INTO mydb.ftdata (entrytime, tank_id, temperature, ph) VALUES (%s,0,0,0)",
        dt))
    if err != nil {
        fmt.Println(err)
    }
}
