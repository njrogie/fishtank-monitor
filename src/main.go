package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize app state
	devicePerformingAction := false
	cmdQueue := []int{}
	router := gin.Default()

	// This is the route that the hardware will periodically check. First duty cycle is 1Hz
	router.GET("/api/currStep", getCurrCmd(&cmdQueue))

	// These are the response methods used to confirm progress during the routines
	router.GET("/api/accept", commandAccepted(&devicePerformingAction, &cmdQueue))
	router.GET("/api/complete", commandCompleted(&devicePerformingAction, &cmdQueue))

	router.GET("/api/test", testAddStep(&cmdQueue))

	router.Run("10.128.0.3:8080")
}
