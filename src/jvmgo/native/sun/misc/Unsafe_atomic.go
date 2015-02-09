package misc

import (
    //"unsafe"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
)

func init() {
    _unsafe(compareAndSwapInt,      "compareAndSwapInt",        "(Ljava/lang/Object;JII)Z")
}

// public final native boolean compareAndSwapInt(Object o, long offset, int expected, int x);
// (Ljava/lang/Object;JII)Z
func compareAndSwapInt(frame *rtda.Frame) {
    stack := frame.OperandStack()
    x := stack.PopInt()
    expected := stack.PopInt()
    offset := stack.PopLong()
    o := stack.PopRef()
    stack.PopRef() // this

    // todo
    fields := o.Fields().([]Any)
    slot := int(offset)
    actual := fields[slot].(int32)
    if actual == expected {
        fields[slot] = x
        stack.PushBoolean(true)
    } else {
        stack.PushBoolean(false)
    }
}
