package misc

import (
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _signal(findSignal, "findSignal",   "(Ljava/lang/String;)I")
    _signal(handle0,    "handle0",      "(IJ)J")
}

func _signal(method Any, name, desc string) {
    rtc.RegisterNativeMethod("sun/misc/Signal", name, desc, method)
}

// private static native int findSignal(String string);
// (Ljava/lang/String;)I
func findSignal(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopRef() // name
    stack.PushInt(0) // todo
}

// private static native long handle0(int i, long l);
// (IJ)J
func handle0(frame *rtda.Frame) {
    stack := frame.OperandStack()
    // todo
    stack.PopLong()
    stack.PopInt()
    stack.PushLong(0)
}

// private static native void raise0(int i);
