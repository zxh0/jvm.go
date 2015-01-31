package lang

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _class(desiredAssertionStatus0, "desiredAssertionStatus0",  "(Ljava/lang/Class;)Z")
    _class(getClassLoader0,         "getClassLoader0",          "()Ljava/lang/ClassLoader;")
    _class(getName0,                "getName0",                 "()Ljava/lang/String;")
    _class(getPrimitiveClass,       "getPrimitiveClass",        "(Ljava/lang/String;)Ljava/lang/Class;")
}

func _class(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/Class", name, desc, method)
}

func desiredAssertionStatus0(frame *rtda.Frame) {
    // todo
    stack := frame.OperandStack()
    _ = stack.PopRef() // this
    stack.PushBoolean(false)
}

func getClassLoader0(frame *rtda.Frame) {
    // todo
    stack := frame.OperandStack()
    _ = stack.PopRef() // this
    stack.PushRef(nil)
}

func getName0(frame *rtda.Frame) {
    panic("getName0")
}

//static native Class<?> getPrimitiveClass(String name);
//getPrimitiveClass(Ljava/lang/String;)Ljava/lang/Class;
func getPrimitiveClass(frame *rtda.Frame) {
    // todo
    panic("getPrimitiveClass")
}
