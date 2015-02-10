package reflect

import (
    //"unsafe"
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _reflection(getCallerClass,         "getCallerClass",       "(I)Ljava/lang/Class;")
    _reflection(getClassAccessFlags,    "getClassAccessFlags",  "(Ljava/lang/Class;)I")
}

func _reflection(method Any, name, desc string) {
    rtc.RegisterNativeMethod("sun/reflect/Reflection", name, desc, method)
}

// public static native Class<?> getCallerClass(int i);
// (I)Ljava/lang/Class;
func getCallerClass(frame *rtda.Frame, x int) {
    vars := frame.LocalVars()
    i := uint(vars.GetInt(0))

    callerClass := frame.Thread().TopFrameN(i).Method().Class().JClass()
    stack := frame.OperandStack()
    stack.PushRef(callerClass)
}

// public static native int getClassAccessFlags(Class<?> type);
// (Ljava/lang/Class;)I
func getClassAccessFlags(frame *rtda.Frame, x int) {
    vars := frame.LocalVars()
    _type := vars.GetRef(0)

    goClass := _type.Extra().(*rtc.Class)
    flags := goClass.GetAccessFlags()

    stack := frame.OperandStack()
    stack.PushInt(int32(flags))
}
