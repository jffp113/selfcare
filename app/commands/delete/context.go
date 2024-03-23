package delete

import (
	"fmt"

	"github.com/jffp113/selfcare/business/config"
	"github.com/urfave/cli/v2"
)

func ContextCommands() *cli.Command {
	return &cli.Command{
		Name:         "context",
		Aliases:      []string{"ctx"},
		Usage:        "",
		UsageText:    "",
		Description:  "Delete a context to be used",
		ArgsUsage:    "",
		Category:     "Basic Commands (Beginner)",
		Flags:        []cli.Flag{},
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action:       contextAction,
		OnUsageError: nil,
	}
}

func contextAction(ctx *cli.Context) error {
	name := ctx.Args().First()

	if name == "" {
		return fmt.Errorf("context name cannot be empty")
	}

	path := ctx.String("db")
	cfg, err := config.ConfigFromFile(path)
	if err != nil {
		return fmt.Errorf("error getting db file: %w", err)
	}

	if name == "" {
		return fmt.Errorf("context name cannot be empty")
	}

	if _, ok := cfg.Contexts[name]; !ok {
		return fmt.Errorf("context does not exist")
	}

	delete(cfg.Contexts, name)

	err = config.StoreConfig(path, cfg)
	if err != nil {
		return err
	}

	fmt.Println("context deleted")

	return nil
}
