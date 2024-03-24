package register

import (
	"fmt"

	"github.com/jffp113/selfcare/app/utils"
	"github.com/urfave/cli/v2"
)

func HolidayCommands() *cli.Command {
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
		Action:       holidayAction,
		OnUsageError: nil,
	}
}

func holidayAction(ctx *cli.Context) error {
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
