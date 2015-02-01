package lang

import (
    "unsafe"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _object(getClass, "getClass", "()Ljava/lang/Class;")
    _object(hashCode, "hashCode", "()I")
}

func _object(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/Object", name, desc, method)
}

// public final native Class<?> getClass();
// ()Ljava/lang/Class;
func getClass(frame *rtda.Frame) {
    stack := frame.OperandStack()
    this := stack.PopRef()
    class := this.Class().JClass()
    stack.PushRef(class)
}

// public native int hashCode();
// ()I
func hashCode(frame *rtda.Frame) {
    stack := frame.OperandStack()
    this := stack.PopRef()
    hash := int32(uintptr(unsafe.Pointer(this)))
    stack.PushInt(hash)
}
