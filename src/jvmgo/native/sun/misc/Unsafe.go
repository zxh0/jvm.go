package misc

import (
    //"unsafe"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _unsafe(addressSize,        "addressSize",      "()I")
    _unsafe(arrayBaseOffset,    "arrayBaseOffset",  "(Ljava/lang/Class;)I")
    _unsafe(arrayIndexScale,    "arrayIndexScale",  "(Ljava/lang/Class;)I")
}

func _unsafe(method Any, name, desc string) {
    rtc.RegisterNativeMethod("sun/misc/Unsafe", name, desc, method)
}

// public native int addressSize();
// ()I
func addressSize(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopRef() // this
    //size := unsafe.Sizeof(int)
    stack.PushInt(0)
}

// public native int arrayBaseOffset(Class<?> type);
// (Ljava/lang/Class;)I
func arrayBaseOffset(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopRef() // this
    stack.PopRef() // type
    stack.PushInt(0) // todo
}

// public native int arrayIndexScale(Class<?> type);
// (Ljava/lang/Class;)I
func arrayIndexScale(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopRef() // this
    stack.PopRef() // type
    stack.PushInt(1) // todo
}
