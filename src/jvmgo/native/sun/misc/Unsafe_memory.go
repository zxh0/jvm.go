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
    _unsafe(putByte,        "putByte",          "(JB)V")
    _unsafe(getByte,        "getByte",          "(J)B")
}

// public native long allocateMemory(long bytes);
// (J)J
func allocateMemory(frame *rtda.Frame) {
    stack := frame.OperandStack()
    bytes := stack.PopLong()
    stack.PopRef() // this

    address := allocate(bytes)
    stack.PushLong(address)
}

// public native void freeMemory(long address);
// (J)V
func freeMemory(frame *rtda.Frame) {
    stack := frame.OperandStack()
    address := stack.PopLong()
    stack.PopRef() // this
    free(address)
}

// public native void putByte(long address, byte b);
// (JB)V
func putByte(frame *rtda.Frame) {
    stack := frame.OperandStack()
    b := stack.PopInt()
    address := stack.PopLong()
    stack.PopRef() // this

    mem := memoryAt(address)
    mem[0] = uint8(b)
}

// public native byte getByte(long address);
// (J)B
func getByte(frame *rtda.Frame) {
    stack := frame.OperandStack()
    address := stack.PopLong()
    stack.PopRef() // this

    mem := memoryAt(address)
    b := mem[0]
    stack.PushInt(int32(b))
}

// public native void putLong(long address, long x);
// (JJ)V
func putLong(frame *rtda.Frame) {
    stack := frame.OperandStack()
    value := stack.PopLong()
    address := stack.PopLong()
    stack.PopRef() // this

    mem := memoryAt(address)
    binary.BigEndian.PutUint64(mem, uint64(value))
}
