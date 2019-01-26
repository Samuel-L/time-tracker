package main

import (
	"time"

	"github.com/Samuel-L/time-tracker/internal/action"
)

func main() {
	timerAction := action.Action{
		ActionType: action.StartTimer,
		Project:    "Time Tracker",
		Timestamp:  time.Now(),
	}

	action.Dispatch(&timerAction)
}
