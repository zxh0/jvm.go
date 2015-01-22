package main

import (
    //"fmt"
    "os"
    //"strings"
    //"jvmgo/classfile"
    //"jvmgo/classpath"
    "jvmgo/cmdline"
)

func main() {
    cmd, err := cmdline.ParseCommand(os.Args)
    if err != nil {
        cmdline.PrintUsage()
        return
    }

    startJVM(cmd)
    // jvm := &JVM{}
    // jvm.cp = cmd.Options().Classpath()

    // className := strings.Replace(cmd.Class(), ".", "/", -1)
    // jvm.startup(className)
}
