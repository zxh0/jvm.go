package main

import (
    "os"
    "jvmgo/cmdline"
    "jvmgo/jvm"
)

func main() {
    cmd, err := cmdline.ParseCommand(os.Args)
    if err != nil {
        cmdline.PrintUsage()
    } else {
        vm := jvm.JVM{}
        vm.Startup(cmd)
    }
}
