/*
    This is the simulator that sends garbage data to the
    server so we can test it out.
*/

package main

import (
    "time"
    "os"
    "fmt"
    "math/rand"
    "encoding/json"
    "net/http"
    "bytes"
    "strings"
)

type TankInfo struct {
    Tank_ID int
    Temperature float32
    Ph float32
}

func main() {
    // Determine whether we use a realtime sim or a sped-up sim
    args := os.Args
    simLoopDuration := 10 * time.Second // 10 seconds
    if len(args) > 1 {
        if(args[1] == "--real" || args[1] == "-r") {
            simLoopDuration = (30 * time.Minute) // End app will only generate a datapoint every 30 minutes hopefully
        }
    }

    // Every loop, we shall send a new dataset
    for {
        // generate json with random values
        fmt.Println("Generating a datapoint.")
        thisSample := TankInfo{
            Tank_ID:        rand.Intn(10),
            Temperature:    (rand.Float32() * 50) + 30,
            Ph:             rand.Float32() * 14,
        }

        // Convert struct to .json
        b, err := json.Marshal(thisSample)
        if err != nil {
            fmt.Println("error:", err)
        }
        fmt.Println("JSON data:", string(b))

        // Send POST with Json data
        dat, _ := os.ReadFile("/app/key")
        res, err := http.Post("http://localhost:8081/newData?key=" + strings.TrimSpace(string(dat)), "application/json", bytes.NewBuffer(b))

        if err != nil {
            fmt.Println("Error sending Post Request:", err)
        } else {
            fmt.Println("Successful POST with code", res.Status)
        }

        // Delay
        time.Sleep(simLoopDuration)
    }
}
