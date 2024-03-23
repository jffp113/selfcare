package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jffp113/selfcare/app/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	dir, err := getDefaultDBDir()
	if err != nil {
		return err
	}

	app := &cli.App{
		Name:  "selfcare",
		Usage: "selcare cli",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "url",
				Value:   "192.168.0.253:6443",
				Aliases: []string{"l"},
				Usage:   "service url",
			},
			/*&cli.StringFlag{
				Name:    "user",
				Value:   "",
				Aliases: []string{"u"},
				Usage:   "service url",
			},
			&cli.StringFlag{
				Name:    "password",
				Value:   "",
				Aliases: []string{"p"},
				Usage:   "service url",
			},*/
			&cli.StringFlag{
				Name:    "db",
				Value:   dir,
				Aliases: []string{"d"},
				Usage:   "service url",
			},
		},
		Commands: commands.GetCommands(),
	}

	ctx, err := createDepsContext()

	if err != nil {
		return fmt.Errorf("creating depedency context: %w", err)
	}

	if err := app.RunContext(ctx, os.Args); err != nil {
		return err
	}

	return nil
}

func createDepsContext() (context.Context, error) {
	//environmentController := environment.New(fmt.Sprintf("https://%v", cfg.Service.Url))
	//context = ctx.SetValue(context, environment.ControllerKey, environmentController)
	return context.Background(), nil
}

func getDefaultDBDir() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}

	return filepath.Join(dir, ".selfcare", "db.yaml"), nil
}
