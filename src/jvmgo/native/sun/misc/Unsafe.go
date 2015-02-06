package misc

import (
    //"unsafe"
    "encoding/binary"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

var _allocatedMemories = map[int64][]byte{}

func init() {
    _unsafe(addressSize,        "addressSize",          "()I")
    _unsafe(allocateMemory,     "allocateMemory",       "(J)J")
    _unsafe(arrayBaseOffset,    "arrayBaseOffset",      "(Ljava/lang/Class;)I")
    _unsafe(arrayIndexScale,    "arrayIndexScale",      "(Ljava/lang/Class;)I")
    _unsafe(compareAndSwapInt,  "compareAndSwapInt",    "(Ljava/lang/Object;JII)Z")
    _unsafe(objectFieldOffset,  "objectFieldOffset",    "(Ljava/lang/reflect/Field;)J")
    _unsafe(putLong,            "putLong",              "(JJ)V")
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

// public native long allocateMemory(long l);
// (J)J
func allocateMemory(frame *rtda.Frame) {
    stack := frame.OperandStack()
    size := stack.PopLong()
    stack.PopRef() // this

    address := allocate(size)
    stack.PushLong(address)
}

// public native int arrayBaseOffset(Class<?> type);
// (Ljava/lang/Class;)I
func arrayBaseOffset(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopRef() // type
    stack.PopRef() // this
    stack.PushInt(0) // todo
}

// public native int arrayIndexScale(Class<?> type);
// (Ljava/lang/Class;)I
func arrayIndexScale(frame *rtda.Frame) {
    stack := frame.OperandStack()
    stack.PopRef() // type
    stack.PopRef() // this
    stack.PushInt(1) // todo
}

// public final native boolean compareAndSwapInt(Object o, long l, int i, int i1);
// (Ljava/lang/Object;JII)Z
func compareAndSwapInt(frame *rtda.Frame) {
    stack := frame.OperandStack()
    i1 := stack.PopInt()
    i := stack.PopInt()
    l := stack.PopLong()
    o := stack.PopRef()
    stack.PopRef() // this

    // todo
    fields := o.Fields().([]Any)
    slot := int(l)
    val := fields[slot].(int32)
    if val == i {
        fields[slot] = i1
        stack.PushBoolean(true)
    } else {
        stack.PushBoolean(false)
    }
}

// public native long objectFieldOffset(Field field);
// (Ljava/lang/reflect/Field;)J
func objectFieldOffset(frame *rtda.Frame) {
    stack := frame.OperandStack()
    jField := stack.PopRef()
    stack.PopRef() // this

    offset := jField.GetFieldValue("slot", "I").(int32)
    stack.PushLong(int64(offset))
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
