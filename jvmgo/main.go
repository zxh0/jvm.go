package main

import (
	"jvmgo/cmdline"
	"jvmgo/jvm"
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
