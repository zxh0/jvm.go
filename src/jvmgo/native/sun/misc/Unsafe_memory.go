package misc

import (
    //"unsafe"
    "encoding/binary"
    "jvmgo/jvm/rtda"
)

func init() {
    _unsafe(allocateMemory, "allocateMemory",   "(J)J")
    _unsafe(freeMemory,     "freeMemory",       "(J)V")
    _unsafe(putLong,        "putLong",          "(JJ)V")
    _unsafe(getByte,        "getByte",          "(J)B")
}

// public native long allocateMemory(long l);
// (J)J
func allocateMemory(frame *rtda.Frame) {
    stack := frame.OperandStack()
    size := stack.PopLong()
    stack.PopRef() // this

    address := allocate(size)
    stack.PushLong(address)
}

// public native void freeMemory(long l);
// (J)V
func freeMemory(frame *rtda.Frame) {
    stack := frame.OperandStack()
    address := stack.PopLong()
    stack.PopRef() // this
    free(address)
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

// public native byte getByte(long l);
// (J)B
func getByte(frame *rtda.Frame) {
    stack := frame.OperandStack()
    address := stack.PopLong()
    stack.PopRef() // this

    mem := memoryAt(address)
    b := mem[0]
    stack.PushInt(int32(b))
}
