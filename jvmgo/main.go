package main

import (
	"github.com/zxh0/jvm.go/jvmgo/cmdline"
	"github.com/zxh0/jvm.go/jvmgo/jvm"
	"os"
)

func main() {
	cmd, err := cmdline.ParseCommand(os.Args)
	if err != nil {
		cmdline.PrintUsage()
	} else {
		jvm.Startup(cmd)
	}
}
