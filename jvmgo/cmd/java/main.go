package main

import (
	"os"
)

func main() {
	cmd, err := ParseCommand(os.Args)
	if err != nil {
		PrintUsage()
	} else {
		startJVM(cmd)
	}
}
