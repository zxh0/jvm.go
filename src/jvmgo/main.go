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
        jvm := JVM{}
        jvm.startup(cmd)
    }
}
