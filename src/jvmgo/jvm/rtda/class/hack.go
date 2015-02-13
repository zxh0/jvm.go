package class

import (
    . "jvmgo/any"
)

// only used by jvm.go
func NewBootstrapMethod(code []byte, classLoader Any) (*Method) {
    method := &Method{}
    method.class = &Class{name:"~jvmgo", classLoader:classLoader.(*ClassLoader)}
    method.name = "<bootstrap>"
    method.accessFlags = ACC_STATIC
    method.maxStack = 8
    method.maxLocals = 8
    method.code = code
    return method
}

func NewGarbageMethod() *Method {
    method := &Method{}
    method.class = &Class{name:"~jvmgo"}
    method.name = "<garbage>"
    method.accessFlags = ACC_STATIC
    method.maxStack = 8
    method.maxLocals = 8
    method.code = []byte{0xb1} // return
    return method
}

// todo
func hackClass(class *Class) {
    if class.name == "java/lang/ClassLoader" {
        loadLibrary := class.GetStaticMethod("loadLibrary", "(Ljava/lang/Class;Ljava/lang/String;Z)V")
        loadLibrary.code = []byte{0xb1} // return void
    }
}
