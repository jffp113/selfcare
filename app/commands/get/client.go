package get

import (
	"fmt"

	"github.com/jffp113/selfcare/app/printer"
	"github.com/jffp113/selfcare/app/utils"
	"github.com/urfave/cli/v2"
)

func ClientCommands() *cli.Command {
	return &cli.Command{
		Name:         "client",
		Aliases:      nil,
		Usage:        "",
		UsageText:    "",
		Description:  "Get user clients",
		ArgsUsage:    "",
		Category:     "Basic Commands (Beginner)",
		Flags:        []cli.Flag{},
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

	clients, err := selfcare.GetClients()
	if err != nil {
		fmt.Println("Error registering timesheet: %w", err)
		return nil
	}

	table := printer.NewTablePrinter("Name", "Id")

	for i := range clients {
		table.AddRow(clients[i].Name, clients[i].Id)
	}

	table.Print()

	return nil
}
