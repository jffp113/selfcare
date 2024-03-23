package add

import (
	"fmt"

	"github.com/jffp113/selfcare/business/config"
	"github.com/urfave/cli/v2"
)

func ContextCommands() *cli.Command {
	return &cli.Command{
		Name:        "context",
		Aliases:     []string{"ctx"},
		Usage:       "",
		UsageText:   "",
		Description: "Add a context to db file",
		ArgsUsage:   "",
		Category:    "Basic Commands (Beginner)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "username",
				Aliases: []string{"u"},
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"p"},
			},
			&cli.StringFlag{
				Name:        "timezone",
				DefaultText: "GMT+0000",
				Aliases:     []string{"t"},
			},
		},
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

	cfg.Contexts[name] = newContext(ctx)

	err = config.StoreConfig(path, cfg)
	if err != nil {
		return err
	}

	fmt.Println("context added")

	return nil
}

func newContext(cliCtx *cli.Context) config.Context {
	var ctx config.Context
	ctx.Username = cliCtx.String("username")
	ctx.Password = cliCtx.String("password")
	ctx.Timezone = cliCtx.String("timezone")
	return ctx
}
