package misc

import (
    //"unsafe"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _unsafe(addressSize,            "addressSize",              "()I")
    _unsafe(arrayBaseOffset,        "arrayBaseOffset",          "(Ljava/lang/Class;)I")
    _unsafe(arrayIndexScale,        "arrayIndexScale",          "(Ljava/lang/Class;)I")
    _unsafe(ensureClassInitialized, "ensureClassInitialized",   "(Ljava/lang/Class;)V")
    _unsafe(objectFieldOffset,      "objectFieldOffset",        "(Ljava/lang/reflect/Field;)J")
    _unsafe(staticFieldOffset,      "staticFieldOffset",        "(Ljava/lang/reflect/Field;)J")
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

// public native void ensureClassInitialized(Class<?> c);
// (Ljava/lang/Class;)V
func ensureClassInitialized(frame *rtda.Frame) {
    stack := frame.OperandStack()
    classObj := stack.PopRef()
    this := stack.PopRef()

    goClass := classObj.Extra().(*rtc.Class)
    if goClass.InitializationNotStarted() {
        // undo ensureClassInitialized()
        frame.RevertNextPC()
        stack.PushRef(this)
        stack.PushRef(classObj)
        // init
        rtda.InitClass(goClass, frame.Thread())
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

// public native long staticFieldOffset(Field field);
// (Ljava/lang/reflect/Field;)J
func staticFieldOffset(frame *rtda.Frame) {
    stack := frame.OperandStack()
    jField := stack.PopRef()
    stack.PopRef() // this

    offset := jField.GetFieldValue("slot", "I").(int32)
    stack.PushLong(int64(offset))
}
