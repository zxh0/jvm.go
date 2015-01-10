package main

import (
    "jvmgo/classpath"
    "jvmgo/heap"
)

type JVM struct {
    cp      *classpath.ClassPath
    heap    *heap.Heap
}
