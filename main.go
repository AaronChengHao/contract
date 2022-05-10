package main

import (
	_ "contract/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"contract/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
