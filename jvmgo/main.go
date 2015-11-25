package main

import (
	"os"

	"github.com/zxh0/jvm.go/jvmgo/cmdline"
)

func main() {
	cmd, err := cmdline.ParseCommand(os.Args)
	if err != nil {
		cmdline.PrintUsage()
	} else {
		startJVM(cmd)
	}
}
