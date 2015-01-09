package main

import "fmt"
// import "io/ioutil"
import "os"
import "strings"
import "jvmgo/classfile"
import "jvmgo/classpath"

func main() {
    mainClassName := os.Args[1]
    mainClassFileName := strings.Replace(mainClassName, ".", "/", -1) + ".class"
    //fmt.Println(mainClassFileName)

    cp := classpath.ParseClassPath(".")
    data, err := cp.ReadClassData(mainClassFileName)

    if err == nil {
        cr := classfile.NewClassReader(data)
        cf, err := classfile.ParseClassFile(cr)
        fmt.Printf("err: %v \n", err)
        fmt.Printf("cf: %v \n", cf)
    }
}
