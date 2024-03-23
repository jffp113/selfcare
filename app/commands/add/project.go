package add

import "github.com/urfave/cli/v2"

func ProjectCommands() *cli.Command {
	return &cli.Command{
		Name:        "project",
		Aliases:     nil,
		Usage:       "",
		UsageText:   "",
		Description: "Register selfcare timeseet",
		ArgsUsage:   "",
		Category:    "Basic Commands (Beginner)",
		/*Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "value",
				Aliases: []string{"v"},
			},
		},*/
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action:       projectAction,
		OnUsageError: nil,
	}
}

func projectAction(ctx *cli.Context) error {
	/*ctr, err := utils.GetEngineController(ctx)
	if err != nil {
		return err
	}

	if ctx.Args().Len() != 1 {
		return errors.New("orchestration name should be specified")
	}

	inputs, err := toInputs(ctx.StringSlice("value"))
	if err != nil {
		return err
	}

	err = ctr.StartOrchestration(ctx.Context, ctx.Args().First(), inputs...)
	if err != nil {
		return fmt.Errorf("starting orchestration: %w", err)
	}
	*/
	return nil
}
