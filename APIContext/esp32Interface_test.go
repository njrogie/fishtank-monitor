package fishtankmonitor_test

import (
	"testing"

	APIContext "nrogie.com/fishtank-api/APIContext"
)

func TestGetCurrentFunctionCode(t *testing.T) {
	const VERIFY = 1

	testContext := APIContext.Default()
	testContext.QueueCmd(APIContext.ESP32Command{
		Code: VERIFY,
		Name: "FLASH",
	})

	lastCmd := testContext.CurrentCommand()

	if lastCmd.Code != VERIFY {
		t.Errorf("Code not correctly saved to context (have %d, want %d)", lastCmd.Code, VERIFY)
	}
}

func TestReceiveCmdAcknowledge(t *testing.T) {
	testContext := APIContext.Default()

	testContext.BeginQueuedCommand()

	if testContext.IsCommandInProgress() {
		t.Errorf("Status is 'running' but no commands queued")
	}

	testContext.QueueCmd(APIContext.ESP32Command{
		Code: 1,
		Name: "FLASH",
	})

	if testContext.IsCommandInProgress() {
		t.Errorf("Status is 'running' but command not yet started")
	}

	testContext.BeginQueuedCommand()

	if !testContext.IsCommandInProgress() {
		t.Errorf("Status is 'not running' but command has been started")
	}
}

func TestReceiveCmdComplete(t *testing.T) {
	testContext := APIContext.Default()
	testContext.QueueCmd(APIContext.ESP32Command{
		Code: 1,
		Name: "FLASH",
	})

	if testContext.NumCommandsQueued() != 1 {
		t.Errorf("Command not queued.")
	}

	testContext.BeginQueuedCommand()

	if !testContext.IsCommandInProgress() {
		t.Errorf("Status is 'not running' but command has been started")
	}

	testContext.ResolveAndRemoveCommand()

	if testContext.IsCommandInProgress() {
		t.Errorf("Context reporting command in progress when should be resolved")
	}

	if testContext.NumCommandsQueued() > 0 {
		t.Errorf("Completed command not removed correctly")
	}

}
