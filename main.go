package main

import (
	"fmt"
	"os"

	"github.com/ikaven1024/container-cli/cli/image"
	"github.com/ikaven1024/container-cli/version"

	"github.com/containerd/containerd/server"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "client for containerd"
	app.Version = version.Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "address",
			Usage: "address of containerd grpc",
			Value: server.DefaultAddress,
		},
	}
	app.Commands = []cli.Command{
		image.Command,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
