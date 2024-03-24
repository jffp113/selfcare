package get

import "github.com/urfave/cli/v2"

func Commands() *cli.Command {
	return &cli.Command{
		Name:         "get",
		Aliases:      nil,
		Usage:        "",
		UsageText:    "",
		Description:  "Get a resource",
		ArgsUsage:    "",
		Category:     "Basic Commands (Beginner)",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action:       nil,
		OnUsageError: nil,
		Subcommands: cli.Commands{
			ClientCommands(),
		},
	}
}
