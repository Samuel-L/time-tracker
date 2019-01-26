package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Samuel-L/time-tracker/internal/action"
	"github.com/Samuel-L/time-tracker/internal/helpers"
	"github.com/joho/godotenv"

	"github.com/urfave/cli"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := cli.NewApp()
	app.Name = "tt (time tracker)"
	app.Description = "Track your time working on projects!"

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start tracking a project",
			Action: func(c *cli.Context) error {
				return start(c)
			},
		},
		{
			Name:  "stop",
			Usage: "stop tracking a project",
			Action: func(c *cli.Context) error {
				return stop(c)
			},
		},
		{
			Name:  "day",
			Usage: "view your day",
			Action: func(c *cli.Context) error {
				return day(c)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func start(ctx *cli.Context) error {
	project := ctx.Args().First()
	if project == "" {
		return cli.NewExitError("Usage: tt start <project_name>", 2)
	}
	startAction := action.Action{
		ActionType: action.StartTimer,
		Project:    project,
		Timestamp:  time.Now(),
	}
	err := action.Dispatch(&startAction)
	if err != nil {
		return err
	}
	return nil
}

func stop(ctx *cli.Context) error {
	project := ctx.Args().First()
	if project == "" {
		return cli.NewExitError("Usage: tt stop <project_name>", 2)
	}
	stopAction := action.Action{
		ActionType: action.StopTimer,
		Project:    project,
		Timestamp:  time.Now(),
	}
	err := action.Dispatch(&stopAction)
	if err != nil {
		return err
	}
	return nil
}

func day(ctx *cli.Context) error {
	data, err := action.Fetch()
	if err != nil {
		return err
	}

	day := []action.Action{}
	for _, value := range data {
		if helpers.IsToday(value.Timestamp) {
			day = append(day, value)
		}
	}

	var previousAction action.Action
	tracked := make(map[string]time.Duration)
	for _, value := range day {
		if previousAction == (action.Action{}) {
			previousAction = value
			continue
		}

		if previousAction.ActionType == action.StartTimer && value.ActionType == action.StopTimer {
			duration := value.Timestamp.Sub(previousAction.Timestamp)
			tracked[value.Project] = duration
		}

		previousAction = value
	}

	for project, duration := range tracked {
		fmt.Printf("%s: %s\n", project, duration)
	}

	return nil
}
