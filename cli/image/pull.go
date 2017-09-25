package image

import (
	//"context"
	"fmt"

	//"github.com/ikaven1024/container-cli/util"

	//"github.com/containerd/containerd"
	"github.com/urfave/cli"
)

var (
	pullCommand = cli.Command{
		Name:   "pull",
		Usage:  "pull image",
		Action: pull,
	}
)

func pull(ctx *cli.Context) error {
	//client, err := util.ContainerClientFromCliContext(ctx)
	//if err != nil {
	//	return nil
	//}

	ref := ctx.Args().First()
	fmt.Printf("pull image %v\n", ref)
	//client.Pull(context.Background(), ref)
	return nil
}
