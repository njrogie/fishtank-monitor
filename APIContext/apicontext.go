package apicontext

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type APIContext struct {
	queue             []ESP32Command
	commandInProgress bool
}

type ESP32Command struct {
	Name string
	Code int
}

func Default() APIContext {
	handler := APIContext{[]ESP32Command{}, false}
	return handler
}

func NewCommand() ESP32Command {
	return ESP32Command{
		Name: "NULL",
		Code: -1,
	}
}

func (ctx *APIContext) NumCommandsQueued() int {
	return len(ctx.queue)
}

func (ctx *APIContext) IsCommandInProgress() bool {
	return ctx.commandInProgress
}

func (ctx *APIContext) QueueCmd(cmd ESP32Command) {
	ctx.queue = append(ctx.queue, cmd)
}

func (ctx *APIContext) BeginQueuedCommand() {
	if ctx.NumCommandsQueued() > 0 {
		ctx.commandInProgress = true
	}
}

func (ctx *APIContext) CurrentCommand() ESP32Command {
	data := NewCommand()
	if ctx.NumCommandsQueued() != 0 {
		data = ctx.queue[0]
	}
	return data
}

func (ctx *APIContext) ResolveAndRemoveCommand() {
	ctx.commandInProgress = false
	if len(ctx.queue) > 1 {
		ctx.queue = ctx.queue[1:]
	} else {
		ctx.queue = []ESP32Command{}
	}
}

func commandCompleted(handler *APIContext) gin.HandlerFunc {
	handlerFn := func(c *gin.Context) {
		if handler.NumCommandsQueued() > 0 {
			/*
				handler.PerformingAction = false
				fmt.Printf("Device is done with command: %d\n", handler.Queue[0])

				if len(handler.Queue) > 1 {
					handler.Queue = (handler.Queue)[1:]
				} else {
					handler.Queue = []int{}
				}
			*/
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
