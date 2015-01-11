package main

import (
    "fmt"
    "jvmgo/classfile"
    "jvmgo/classpath"
    "jvmgo/rtda"
)

type JVM struct {
    cp          *classpath.ClassPath
    methodArea  *rtda.MethodArea
    heap        *rtda.Heap
}

func (self *JVM) startup(className string) {
    self.heap = rtda.NewHeap()
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
    
    class := rtda.NewClass(cf)
    fmt.Printf("class: %v \n", class)
}
