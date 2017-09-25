package util

import (
	"github.com/containerd/containerd"
	"github.com/urfave/cli"
)

func ContainerClientFromCliContext(ctx *cli.Context) (*containerd.Client, error) {
	return containerd.New(ctx.GlobalString("address"),
		containerd.WithDefaultNamespace(ctx.GlobalString("namespace")))
}
