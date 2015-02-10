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
    vars := frame.LocalVars()
    // vars.GetRef(0) // this
    o := vars.GetRef(1)
    offset := vars.GetLong(2)
    expected := vars.GetInt(4)
    x := vars.GetInt(5)

    // todo
    fields := o.Fields().([]Any)
    slot := int(offset)
    actual := fields[slot].(int32)
    stack := frame.OperandStack()
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
    vars := frame.LocalVars()
    // vars.GetRef(0) // this
    o := vars.GetRef(1)
    offset := vars.GetLong(2)
    expected := vars.GetLong(4)
    x := vars.GetLong(6)

    // todo
    fields := o.Fields().([]Any)
    slot := int(offset)
    actual := fields[slot].(int64)
    stack := frame.OperandStack()
    if actual == expected {
        fields[slot] = x
        stack.PushBoolean(true)
    } else {
        stack.PushBoolean(false)
    }
}
