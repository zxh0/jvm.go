package main

import (
    "os"
    "jvmgo/cmdline"
)

func main() {
    cmd, err := cmdline.ParseCommand(os.Args)
    if err != nil {
        cmdline.PrintUsage()
    } else {
        startJVM(cmd)
    }
}
