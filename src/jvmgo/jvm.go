package main

import (
    "jvmgo/classpath"
    "jvmgo/heap"
)

type JVM struct {
    cp      *classpath.ClassPath
    heap    *heap.Heap
}

func (self *JVM) startup(className string) {
    self.heap = heap.NewHeap()
    // todo
    // load class
    // find main method
    // execute main
}
