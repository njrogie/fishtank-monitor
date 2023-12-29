package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getState(app *appState) gin.HandlerFunc {
	handlerFn := func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, *app)
	}
	return handlerFn
}

func getCurrCmd(queue *[]int) gin.HandlerFunc {
	handlerFn := func(c *gin.Context) {
		data := -1
		if len(*queue) != 0 {
			data = (*queue)[0]
		}
		c.IndentedJSON(http.StatusOK, data)
	}

	return handlerFn
}

func commandAccepted(actionStatus *bool, queue *[]int) gin.HandlerFunc {
	handlerFn := func(c *gin.Context) {
		if len(*queue) > 0 {
			*actionStatus = true
			fmt.Printf("Device is performing command: %d\n", (*queue)[0])
		}
	}
	return handlerFn
}

func commandCompleted(actionStatus *bool, queue *[]int) gin.HandlerFunc {
	handlerFn := func(c *gin.Context) {
		if len(*queue) > 0 {
			*actionStatus = false
			fmt.Printf("Device is done with command: %d\n", (*queue)[0])

			if len(*queue) > 1 {
				*queue = (*queue)[1:]
			} else {
				*queue = []int{}
			}
		}
	}
	return handlerFn
}

func testAddStep(queue *[]int) gin.HandlerFunc {
	handlerFn := func(c *gin.Context) {
		*queue = append(*queue, 1)
		fmt.Println(*queue)
	}
	return handlerFn
}
