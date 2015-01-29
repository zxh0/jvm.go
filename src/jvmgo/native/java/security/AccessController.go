package security

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

func init() {
    _ac("doPrivileged", "(Ljava/security/PrivilegedAction;)Ljava/lang/Object;", doPrivileged)
}

func _ac(name, desc string, method Any) {
    rtc.RegisterNativeMethod("java/security/AccessController", name, desc, method)
}

//doPrivileged(Ljava/security/PrivilegedAction;)Ljava/lang/Object;
//public static native <T> T doPrivileged(PrivilegedAction<T> action);
func doPrivileged(operandStack *rtda.OperandStack) {
    // todo
    panic("doPrivileged")
}
