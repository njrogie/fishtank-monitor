package main

type appState struct {
	Motors        []motorState `json:"motors"`
	CurrRoutine   string       `json:"routine"`
	RoutineActive bool         `json:"routine_active"`
}

type motorState struct {
	MotorID    int    `json:"motor_id"`
	MotorTitle string `json:"motor_title"`
	Enabled    bool   `json:"enabled"`
}

func initializeAppState() appState {
	// List out the motors
	motors := []motorState{
		{
			MotorID:    0,
			MotorTitle: "Chem1",
			Enabled:    false,
		},
		{
			MotorID:    1,
			MotorTitle: "Chem2",
			Enabled:    false,
		},
		{
			MotorID:    2,
			MotorTitle: "Tank",
			Enabled:    false,
		},
		{
			MotorID:    3,
			MotorTitle: "Drain",
			Enabled:    false,
		},
	}
	// Initialize the state
	return appState{
		Motors:        motors,
		RoutineActive: false,
	}
}
