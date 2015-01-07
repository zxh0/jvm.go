package main

import "fmt"
import "io/ioutil"
import "os"
import "strings"
import "jvmgo/classfile"

func main() {
    mainClassName := os.Args[1]
    mainClassFileName := strings.Replace(mainClassName, ".", "/", -1) + ".class"
    //fmt.Println(mainClassFileName)

    mainClassFile, err := os.Open(mainClassFileName)
    if err != nil {
        fmt.Println("can not open file!")
        fmt.Println(err.Error())
    } else {
        defer mainClassFile.Close()
        fmt.Println("ok")
        readFile(mainClassFileName)
    }
}

// todo
func readFile(file string) {
    bytes, err := ioutil.ReadFile(file) // func ReadFile(filename string) ([]byte, error)
    if err != nil {
        fmt.Println("read file err")
    } else {
        fmt.Println("read file ok")
        fmt.Println(len(bytes))

        // todo
        cr := classfile.NewClassReader(bytes)
        cf, err := classfile.ParseClassFile(cr)
        fmt.Printf("err: %v \n", err)
        fmt.Printf("cf: ", cf)
    }
}
