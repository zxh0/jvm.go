package security

import (
    //"fmt"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
    //_ "jvmgo/instructions"
)

func init() {
    _ac(doPrivileged, "doPrivileged", "(Ljava/security/PrivilegedAction;)Ljava/lang/Object;")
}

func _ac(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/security/AccessController", name, desc, method)
}

//doPrivileged(Ljava/security/PrivilegedAction;)Ljava/lang/Object;
//public static native <T> T doPrivileged(PrivilegedAction<T> action);
func doPrivileged(frame *rtda.Frame) {
    stack := frame.OperandStack()
    action := stack.PopRef()
    methodref := action.Class().ConstantPool().GetMethodref("run") // todo
    method := methodref.VirtualMethod(action)

    stack.PushRef(action)
    frame.Thread().InvokeMethod(method)
}
