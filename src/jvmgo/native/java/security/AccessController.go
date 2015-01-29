package security

import (
    "fmt"
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
    mref := action.Class().ConstantPool().GetMethodref("run") // todo
    fmt.Printf("mref::%v\n", mref)

    // todo
    panic("doPrivileged")
}
