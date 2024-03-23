package utils

import (
	"github.com/jffp113/selfcare/business/controller/selfcare"
	"github.com/urfave/cli/v2"
)

func GetSelfcare(ctx *cli.Context) (*selfcare.Selfcare, error) {
	return selfcare.New(
		selfcare.WithConfigFilePath(ctx.String("db")),
		selfcare.WithBaseUrl(ctx.String("url")),
	)
}
