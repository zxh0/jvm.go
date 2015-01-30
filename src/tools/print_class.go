package main

import (
    "archive/zip"
    "fmt"
    "os"
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

    // find class
    for _, f := range r.File {
        fmt.Printf("%v\n", f.Name)
        // if f.Name == className {
        //     rc, err := f.Open() // func (f *File) Open() (rc io.ReadCloser, err error)
        //     if err != nil {
        //         return nil, err
        //     }
        //     // read class data
        //     data, err := ioutil.ReadAll(rc) // func ReadAll(r io.Reader) ([]byte, error)
        //     rc.Close()
        //     if err != nil {
        //         return nil, err
        //     }
        //     return data, nil
        // }
    }
}
