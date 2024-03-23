package delete

import "github.com/urfave/cli/v2"

func Commands() *cli.Command {
	return &cli.Command{
		Name:         "delete",
		Aliases:      nil,
		Usage:        "",
		UsageText:    "",
		Description:  "Delete a resource",
		ArgsUsage:    "",
		Category:     "Basic Commands (Beginner)",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action:       nil,
		OnUsageError: nil,
		Subcommands: cli.Commands{
			ContextCommands(),
		},
	}
}
