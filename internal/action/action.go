package action

import (
	"fmt"
	"time"
)

// StartTimer constant
const StartTimer string = "START_TIMER"

// StopTimer constant
const StopTimer string = "STOP_TIMER"

// Action type
type Action struct {
	ActionType string
	Project    string
	Timestamp  time.Time
}

func (action *Action) toString() string {
	return fmt.Sprintf("Action: %s, Project: %s, Timestamp: %s",
		action.ActionType,
		action.Project,
		action.Timestamp.Format(time.RFC3339),
	)
}

// Dispatch an action
func Dispatch(action *Action) {
	fmt.Printf("DISPATCHING ACTION: %s\n", action.toString())
}
