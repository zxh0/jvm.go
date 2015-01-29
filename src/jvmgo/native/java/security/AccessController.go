package security

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
    //_ "jvmgo/instructions"
)

func init() {
    _ac("doPrivileged", "(Ljava/security/PrivilegedAction;)Ljava/lang/Object;", doPrivileged)
}

func _ac(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/security/AccessController", name, desc, method)
}

//doPrivileged(Ljava/security/PrivilegedAction;)Ljava/lang/Object;
//public static native <T> T doPrivileged(PrivilegedAction<T> action);
func doPrivileged(frame *rtda.Frame) {
    stack := frame.OperandStack()
    action := stack.PopRef()
    action.Class()


    // todo
    panic("doPrivileged")
}
