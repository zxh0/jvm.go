package main

import (
    "archive/zip"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "jvmgo/classfile"
)

func main() {
    if len(os.Args) > 1 {
        jarFileName := os.Args[1]
        handleJar(jarFileName)
    }
}

func handleJar(jarFileName string) {
    fmt.Printf("jar: %v\n", jarFileName)

    // open jar
    r, err := zip.OpenReader(jarFileName) // func OpenReader(name string) (*ReadCloser, error)
    if err != nil {
        panic(err.Error())
    }
    defer r.Close()

    // find classes
    for _, f := range r.File {
        if strings.HasSuffix(f.Name, ".class") {
            handleClass(f)
        }
    }
}

func handleClass(f *zip.File) {
    //fmt.Printf("%v\n", f.Name)
    
    // open classfile
    rc, err := f.Open() // func (f *File) Open() (rc io.ReadCloser, err error)
    if err != nil {
        panic(err.Error())
    }

    // read class data
    data, err := ioutil.ReadAll(rc) // func ReadAll(r io.Reader) ([]byte, error)
    rc.Close()
    if err != nil {
        panic(err.Error())
    }

    // parse classfile
    cf, err := classfile.ParseClassFile(data)
    if err != nil {
        panic(err.Error())
    }

    handleClassfile(cf)
}

func handleClassfile(cf *classfile.ClassFile) {
    for _, m := range cf.Methods() {
        if m.IsNative() {
            if m.IsStatic() {
                fmt.Printf("%v.%v%v\n", cf.ClassName(), m.Name(), m.Descriptor())
            } else {
                fmt.Printf("%v#%v%v\n", cf.ClassName(), m.Name(), m.Descriptor())
            }
        }
    }
}
