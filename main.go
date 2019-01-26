package main

import (
	"log"
	"os"
	"time"

	"github.com/Samuel-L/time-tracker/internal/action"
	"github.com/joho/godotenv"

	"github.com/urfave/cli"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .new file")
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
