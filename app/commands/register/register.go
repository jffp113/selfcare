package register

import "github.com/urfave/cli/v2"

func Commands() *cli.Command {
	return &cli.Command{
		Name:         "register",
		Aliases:      nil,
		Usage:        "",
		UsageText:    "",
		Description:  "Register a resource",
		ArgsUsage:    "",
		Category:     "Basic Commands (Beginner)",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action:       nil,
		OnUsageError: nil,
		Subcommands: cli.Commands{
			TimesheetCommands(),
		},
	}
}
