package fishtankmonitor

type APIContext struct {
	queue             []ESP32Command
	commandInProgress bool
}

type ESP32Command struct {
	Name string
	Code int
}

type MotorStatus struct {
	name string
	isOn bool
}

// Generate new APIContext
func Default() APIContext {
	handler := APIContext{[]ESP32Command{}, false}
	return handler
}

// Generate new ESP32 Command
func NewCommand() ESP32Command {
	return ESP32Command{
		Name: "NULL",
		Code: -1,
	}
}

func NewMotor(motorName string) MotorStatus {
	return MotorStatus{
		name: motorName,
		isOn: false,
	}
}

// Return number of commands queued
func (ctx *APIContext) NumCommandsQueued() int {
	return len(ctx.queue)
}

// Return status of command in progress
func (ctx *APIContext) IsCommandInProgress() bool {
	return ctx.commandInProgress
}

// Queue a new ESP32Command
func (ctx *APIContext) QueueCmd(cmd ESP32Command) {
	ctx.queue = append(ctx.queue, cmd)
}

// Begin the last queued command
func (ctx *APIContext) BeginQueuedCommand() {
	if ctx.NumCommandsQueued() > 0 {
		ctx.commandInProgress = true
	}
}

// Retrieve the currently queued command
func (ctx *APIContext) CurrentCommand() ESP32Command {
	data := NewCommand()
	if ctx.NumCommandsQueued() != 0 {
		data = ctx.queue[0]
	}
	return data
}

// Resolve/complete the command, and remove it from the queue
func (ctx *APIContext) ResolveAndRemoveCommand() {
	ctx.commandInProgress = false
	if len(ctx.queue) > 1 {
		ctx.queue = ctx.queue[1:]
	} else {
		ctx.queue = []ESP32Command{}
	}
}
