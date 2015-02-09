package misc

import (
    "encoding/binary"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    "jvmgo/util/bigendian"
)

func init() {
    _unsafe(allocateMemory, "allocateMemory",   "(J)J")
    _unsafe(freeMemory,     "freeMemory",       "(J)V")
    _unsafe(putLong,        "putLong",          "(JJ)V")
    _unsafe(putByte,        "putByte",          "(JB)V")
    _unsafe(getByte,        "getByte",          "(J)B")
    _unsafe(putShort,       "putShort",         "(JS)V")
    _unsafe(getShort,       "getShort",         "(J)S")
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

// public native void putByte(long address, byte x);
// (JB)V
func putByte(frame *rtda.Frame) {
    mem, value := _put(frame)
    bigendian.PutInt8(mem, int8(value.(int32)))
}

// public native byte getByte(long address);
// (J)B
func getByte(frame *rtda.Frame) {
    stack, mem := _get(frame)
    stack.PushInt(int32(bigendian.Int8(mem)))
}

// public native void putShort(long address, short x);
// (JS)V
func putShort(frame *rtda.Frame) {
    mem, value := _put(frame)
    bigendian.PutInt16(mem, int16(value.(int32)))
}

// public native short getShort(long address);
// (J)S
func getShort(frame *rtda.Frame) {
    stack, mem := _get(frame)
    stack.PushInt(int32(bigendian.Int16(mem)))
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

func _put(frame *rtda.Frame) ([]byte, Any) {
    stack := frame.OperandStack()
    value := stack.Pop()
    address := stack.PopLong()
    stack.PopRef() // this

    mem := memoryAt(address)
    return mem, value
}

func _get(frame *rtda.Frame) (*rtda.OperandStack, []byte) {
    stack := frame.OperandStack()
    address := stack.PopLong()
    stack.PopRef() // this

    mem := memoryAt(address)
    return stack, mem
}
