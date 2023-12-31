package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	APIContext "nrogie.com/fishtank-api/APIContext"
)

func main() {
	// load authentication details from dotenv file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading auth data")
	}

	esp32user, esp32pass := os.Getenv("ESP32USER"), os.Getenv("ESP32PASS")
	esp32Auth := gin.BasicAuth(gin.Accounts{esp32user: esp32pass})
	jsUser, jsPass := os.Getenv("JSUSER"), os.Getenv("JSPASS")
	jsAuth := gin.BasicAuth(gin.Accounts{jsUser: jsPass})
	router := gin.Default()

	apiContext := APIContext.Default()

	// ESP32 Comm Endpoints
	router.GET("/api/currCmd", esp32Auth, func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, apiContext.CurrentCommand())
	})
	router.POST("/api/beginCmd", esp32Auth, func(c *gin.Context) {
		apiContext.BeginQueuedCommand()
		c.Status(http.StatusAccepted)
	})
	router.POST("/api/finishedCmd", esp32Auth, func(c *gin.Context) {
		apiContext.ResolveAndRemoveCommand()
		c.Status(http.StatusAccepted)
	})

	// JS Comm Endpoints
	router.GET("/api/status", jsAuth, func(c *gin.Context) {

	})

	router.Run("127.0.0.1:8080")
}
