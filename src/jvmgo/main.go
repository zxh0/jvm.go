package main

import (
    "fmt"
    "os"
    "strings"
    "jvmgo/classfile"
    //"jvmgo/classpath"
    "jvmgo/cmdline"
)

func main() {
    cmd, err := cmdline.ParseCommand(os.Args)
    if err != nil {
        cmdline.PrintUsage()
        return
    }

    cp := cmd.Options().Classpath()
    className := strings.Replace(cmd.Class(), ".", "/", -1)

    data, err := cp.ReadClassData(className)

    if err == nil {
        cr := classfile.NewClassReader(data)
        cf, err := classfile.ParseClassFile(cr)
        fmt.Printf("err: %v \n", err)
        fmt.Printf("cf: %v \n", cf)
    } else {
        fmt.Printf("err: %v \n", err)
    }
}
