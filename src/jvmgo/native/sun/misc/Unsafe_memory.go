package misc

import (
    //"unsafe"
    "encoding/binary"
    "jvmgo/jvm/rtda"
)

// public native long allocateMemory(long l);
// (J)J
func allocateMemory(frame *rtda.Frame) {
    stack := frame.OperandStack()
    size := stack.PopLong()
    stack.PopRef() // this

    address := allocate(size)
    stack.PushLong(address)
}

// public native void putLong(long l, long l1);
// (JJ)V
func putLong(frame *rtda.Frame) {
    stack := frame.OperandStack()
    value := stack.PopLong()
    address := stack.PopLong()
    stack.PopRef() // this

    mem := memoryAt(address)
    binary.BigEndian.PutUint64(mem, uint64(value))
}
