package main

import (
    "fmt"
    "jvmgo/classfile"
    "jvmgo/classpath"
    "jvmgo/heap"
)

type JVM struct {
    cp      *classpath.ClassPath
    heap    *heap.Heap
}

func (self *JVM) startup(className string) {
    self.heap = heap.NewHeap()
    self.loadClass(className)
    // todo
    // load class
    // find main method
    // execute main
}

func (self *JVM) loadClass(className string) {
    data, err := self.cp.ReadClassData(className)
    if err != nil {
        // todo
        panic(err.Error())
    }
    cr := classfile.NewClassReader(data)
    cf, err := classfile.ParseClassFile(cr)
    if err != nil {
        // todo
        panic(err.Error())
    }
    fmt.Printf("cf: %v \n", cf)
}
