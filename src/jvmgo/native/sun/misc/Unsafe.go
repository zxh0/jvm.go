package misc

import (
    //"unsafe"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

var _allocatedMemories = map[int64][]byte{}

func init() {
    _unsafe(addressSize,        "addressSize",          "()I")
    _unsafe(arrayBaseOffset,    "arrayBaseOffset",      "(Ljava/lang/Class;)I")
    _unsafe(arrayIndexScale,    "arrayIndexScale",      "(Ljava/lang/Class;)I")
    _unsafe(compareAndSwapInt,  "compareAndSwapInt",    "(Ljava/lang/Object;JII)Z")
    _unsafe(objectFieldOffset,  "objectFieldOffset",    "(Ljava/lang/reflect/Field;)J")
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

// public native void ensureClassInitialized(Class<?> c);
// (Ljava/lang/Class;)V
func ensureClassInitialized(frame *rtda.Frame) {
    stack := frame.OperandStack()
    classObj := stack.PopRef()
    this := stack.PopRef()

    goClass := classObj.Extra().(*rtc.Class)
    if goClass.InitializationNotStarted() {
        thread := frame.Thread()
        // undo ensureClassInitialized()
        stack.PushRef(this)
        stack.PushRef(classObj)
        frame.SetNextPC(thread.PC())
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
