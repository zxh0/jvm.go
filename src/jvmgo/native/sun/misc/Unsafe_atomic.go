package misc

import (
    //"unsafe"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
)

func init() {
    _unsafe(compareAndSwapInt,  "compareAndSwapInt",    "(Ljava/lang/Object;JII)Z")
    _unsafe(compareAndSwapLong, "compareAndSwapLong",   "(Ljava/lang/Object;JJJ)Z")
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

// public final native boolean compareAndSwapLong(Object o, long offset, long expected, long x);
// (Ljava/lang/Object;JJJ)Z
func compareAndSwapLong(frame *rtda.Frame) {
    stack := frame.OperandStack()
    x := stack.PopLong()
    expected := stack.PopLong()
    offset := stack.PopLong()
    o := stack.PopRef()
    stack.PopRef() // this

    // todo
    fields := o.Fields().([]Any)
    slot := int(offset)
    actual := fields[slot].(int64)
    if actual == expected {
        fields[slot] = x
        stack.PushBoolean(true)
    } else {
        stack.PushBoolean(false)
    }
}
