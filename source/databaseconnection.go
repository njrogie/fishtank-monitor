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
