package register

import (
	"fmt"
	"os"

	"github.com/jffp113/selfcare/app/utils"
	"github.com/jffp113/selfcare/business/config"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func TimesheetCommands() *cli.Command {
	return &cli.Command{
		Name:        "timesheet",
		Aliases:     nil,
		Usage:       "",
		UsageText:   "",
		Description: "Register selfcare timeseet",
		ArgsUsage:   "",
		Category:    "Basic Commands (Beginner)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
			},
		},
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action:       timesheetAction,
		OnUsageError: nil,
	}
}

func timesheetAction(ctx *cli.Context) error {
	selfcare, err := utils.GetSelfcare(ctx)
	if err != nil {
		return err
	}

	t, err := getTimesheet(ctx.String("file"))
	if err != nil {
		return err
	}

	err = selfcare.RegisterTimesheet(t)
	if err != nil {
		fmt.Println("Error registering timesheet: %w", err)
		return nil
	}

	fmt.Println("Registed timesheet")

	return nil
}

func getTimesheet(path string) (config.Timesheet, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return config.Timesheet{}, fmt.Errorf("reading timesheet file: %w", err)
	}

	var t config.Timesheet

	err = yaml.Unmarshal(bs, &t)
	if err != nil {
		return config.Timesheet{}, err
	}

	return t, nil
}
