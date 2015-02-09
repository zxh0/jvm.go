package misc

import (
    //"unsafe"
    //. "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _unsafe(ensureClassInitialized, "ensureClassInitialized",   "(Ljava/lang/Class;)V")
    _unsafe(staticFieldOffset,      "staticFieldOffset",        "(Ljava/lang/reflect/Field;)J")
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

// public native long staticFieldOffset(Field field);
// (Ljava/lang/reflect/Field;)J
func staticFieldOffset(frame *rtda.Frame) {
    stack := frame.OperandStack()
    fieldObj := stack.PopRef()
    stack.PopRef() // this

    offset := fieldObj.GetFieldValue("slot", "I").(int32)
    stack.PushLong(int64(offset))
}
