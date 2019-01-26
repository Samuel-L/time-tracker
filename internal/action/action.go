package action

import (
	"fmt"
	"time"

	"github.com/Samuel-L/time-tracker/internal/helpers"
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

type payload struct {
	ActionType string `json:"action"`
	Project    string `json:"project"`
	Timestamp  string `json:"timestamp"`
}

// Dispatch an action
func Dispatch(action *Action) error {
	client, ctx := helpers.FirebaseClient()
	ref := client.NewRef("actions")

	timestamp := action.Timestamp.Format(time.RFC3339)

	if _, err := ref.Push(ctx, &payload{
		ActionType: action.ActionType,
		Project:    action.Project,
		Timestamp:  timestamp,
	}); err != nil {
		return err
	}
	return nil
}
