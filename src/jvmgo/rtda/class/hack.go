package class

import (
    . "jvmgo/any"
)

// only used by exec_main.go
func NewStringArray(strs []*Obj, classLoader *ClassLoader) (*Obj) {
    componentClass := classLoader.StringClass()
    arrClass := classLoader.getRefArrayClass(componentClass)
    return &Obj{arrClass, strs, nil}
}

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
