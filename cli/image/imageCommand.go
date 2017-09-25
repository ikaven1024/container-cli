package image

import (
	"github.com/urfave/cli"
)

var (
	Command = cli.Command{
		Name:  "image",
		Usage: "image operations",
		Subcommands: []cli.Command{
			pullCommand,
		},
	}
)
