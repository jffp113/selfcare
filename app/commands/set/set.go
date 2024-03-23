package set

import "github.com/urfave/cli/v2"

func Commands() *cli.Command {
	return &cli.Command{
		Name:         "set",
		Aliases:      nil,
		Usage:        "",
		UsageText:    "",
		Description:  "Set a resource",
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
