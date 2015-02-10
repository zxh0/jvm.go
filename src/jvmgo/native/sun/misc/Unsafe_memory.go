package misc

import (
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    "jvmgo/util/bigendian"
)

func init() {
    _unsafe(allocateMemory,     "allocateMemory",   "(J)J")
    _unsafe(reallocateMemory,   "reallocateMemory", "(JJ)J")
    _unsafe(freeMemory,         "freeMemory",       "(J)V")
    _unsafe(addressSize,        "addressSize",      "()I")
    _unsafe(putAddress,         "putAddress",       "(JJ)V")
    _unsafe(getAddress,         "getAddress",       "(J)J")
    _unsafe(putByte,            "putByte",          "(JB)V")
    _unsafe(getByte,            "getByte",          "(J)B")
    _unsafe(putShort,           "putShort",         "(JS)V")
    _unsafe(getShort,           "getShort",         "(J)S")
    _unsafe(putChar,            "putChar",          "(JC)V")
    _unsafe(getChar,            "getChar",          "(J)C")
    _unsafe(putInt,             "putInt",           "(JI)V")
    _unsafe(getInt,             "getInt",           "(J)I")
    _unsafe(putLong,            "putLong",          "(JJ)V")
    _unsafe(getLong,            "getLong",          "(J)J")
    _unsafe(putFloat,           "putFloat",         "(JF)V")
    _unsafe(getFloat,           "getFloat",         "(J)F")
    _unsafe(putDouble,          "putDouble",        "(JD)V")
    _unsafe(getDouble,          "getDouble",        "(J)D")
}

// public native long allocateMemory(long bytes);
// (J)J
func allocateMemory(frame *rtda.Frame) {
    vars := frame.LocalVars()
    // vars.GetRef(0) // this
    bytes := vars.GetLong(1)

    address := allocate(bytes)
    stack := frame.OperandStack()
    stack.PushLong(address)
}

// public native long reallocateMemory(long address, long bytes);
// (JJ)J
func reallocateMemory(frame *rtda.Frame) {
    vars := frame.LocalVars()
    // vars.GetRef(0) // this
    address := vars.GetLong(1)
    bytes := vars.GetLong(3)

    newAddress := reallocate(address, bytes)
    stack := frame.OperandStack()
    stack.PushLong(newAddress)
}

// public native void freeMemory(long address);
// (J)V
func freeMemory(frame *rtda.Frame) {
    vars := frame.LocalVars()
    // vars.GetRef(0) // this
    address := vars.GetLong(1)
    free(address)
}

// public native int addressSize();
// ()I
func addressSize(frame *rtda.Frame) {
    // vars := frame.LocalVars()
    // vars.GetRef(0) // this

    stack := frame.OperandStack()
    stack.PushInt(8) // todo unsafe.Sizeof(int)
}

// public native void putAddress(long address, long x);
// (JJ)V
func putAddress(frame *rtda.Frame) {
    putLong(frame)
}

// public native long getAddress(long address);
// (J)J
func getAddress(frame *rtda.Frame) {
    getLong(frame)
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

// public native void putChar(long address, char x);
// (JC)V
func putChar(frame *rtda.Frame) {
    mem, value := _put(frame)
    bigendian.PutUint16(mem, uint16(value.(int32)))
}

// public native char getChar(long address);
// (J)C
func getChar(frame *rtda.Frame) {
    stack, mem := _get(frame)
    stack.PushInt(int32(bigendian.Uint16(mem)))
}

// public native void putInt(long address, int x);
// (JI)V
func putInt(frame *rtda.Frame) {
    mem, value := _put(frame)
    bigendian.PutInt32(mem, value.(int32))
}

// public native int getInt(long address);
// (J)I
func getInt(frame *rtda.Frame) {
    stack, mem := _get(frame)
    stack.PushInt(bigendian.Int32(mem))
}

// public native void putLong(long address, long x);
// (JJ)V
func putLong(frame *rtda.Frame) {
    mem, value := _put(frame)
    bigendian.PutInt64(mem, value.(int64))
}

// public native long getLong(long address);
// (J)J
func getLong(frame *rtda.Frame) {
    stack, mem := _get(frame)
    stack.PushLong(bigendian.Int64(mem))
}

// public native void putFloat(long address, float x);
// (JJ)V
func putFloat(frame *rtda.Frame) {
    mem, value := _put(frame)
    bigendian.PutFloat32(mem, value.(float32))
}

// public native float getFloat(long address);
// (J)J
func getFloat(frame *rtda.Frame) {
    stack, mem := _get(frame)
    stack.PushFloat(bigendian.Float32(mem))
}

// public native void putDouble(long address, double x);
// (JJ)V
func putDouble(frame *rtda.Frame) {
    mem, value := _put(frame)
    bigendian.PutFloat64(mem, value.(float64))
}

// public native double getDouble(long address);
// (J)J
func getDouble(frame *rtda.Frame) {
    stack, mem := _get(frame)
    stack.PushDouble(bigendian.Float64(mem))
}

func _put(frame *rtda.Frame) ([]byte, Any) {
    vars := frame.LocalVars()
    // vars.GetRef(0) // this
    address := vars.GetLong(1)
    value := vars.Get(3)

    mem := memoryAt(address)
    return mem, value
}

func _get(frame *rtda.Frame) (*rtda.OperandStack, []byte) {
    vars := frame.LocalVars()
    // vars.GetRef(0) // this
    address := vars.GetLong(1)

    stack := frame.OperandStack()
    mem := memoryAt(address)
    return stack, mem
}
