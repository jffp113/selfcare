package commands

import (
	"github.com/jffp113/selfcare/app/commands/add"
	"github.com/jffp113/selfcare/app/commands/register"
	"github.com/urfave/cli/v2"
)

func GetCommands() cli.Commands {
	return cli.Commands{
		register.Commands(),
		add.Commands(),
	}
}
