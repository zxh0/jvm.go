package class

import (
    . "jvmgo/any"
    cf "jvmgo/classfile"
)

// only used by string_helper.go
func NewIntArray(ints []int32) (*Obj) {
    return &Obj{nil, ints, nil}
}

// only used by jvm.go
func NewStartupMethod(code []byte, classLoader Any) (*Method) {
    method := &Method{}
    method.class = &Class{name:"~jvmgo", classLoader:classLoader.(*ClassLoader)}
    method.name = "<jvmgo>"
    method.maxStack = 8
    method.maxLocals = 8
    method.code = code
    return method
}

func NewGarbageMethod() *Method {
    method := &Method{}
    method.class = &Class{name:"~garbage"}
    method.name = "<garbage>"
    method.SetAccessFlags(cf.ACC_STATIC)
    method.maxStack = 8
    method.maxLocals = 8
    method.code = []byte{0xb1} // return
    return method
}
