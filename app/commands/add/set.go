package add

import "github.com/urfave/cli/v2"

func Commands() *cli.Command {
	return &cli.Command{
		Name:         "add",
		Aliases:      nil,
		Usage:        "",
		UsageText:    "",
		Description:  "Run a resource",
		ArgsUsage:    "",
		Category:     "Basic Commands (Beginner)",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action:       nil,
		OnUsageError: nil,
		Subcommands: cli.Commands{
			ProjectCommands(),
		},
	}
}
