package security

import (
	//"fmt"
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
	//_ "jvmgo/jvm/instructions"
)

func init() {
	_ac(doPrivileged, "doPrivileged", "(Ljava/security/PrivilegedAction;)Ljava/lang/Object;")
	_ac(getStackAccessControlContext, "getStackAccessControlContext", "()Ljava/security/AccessControlContext;")
}

func _ac(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/security/AccessController", name, desc, method)
}

// public static native <T> T doPrivileged(PrivilegedAction<T> action);
// (Ljava/security/PrivilegedAction;)Ljava/lang/Object;
func doPrivileged(frame *rtda.Frame) {
	vars := frame.LocalVars()
	action := vars.GetRef(0)

	methodref := action.Class().ConstantPool().GetMethodref("run") // todo
	method := methodref.FindInterfaceMethod(action)

	stack := frame.OperandStack()
	stack.PushRef(action)
	frame.Thread().InvokeMethod(method)
}

// private static native AccessControlContext getStackAccessControlContext();
// ()Ljava/security/AccessControlContext;
func getStackAccessControlContext(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushRef(nil)
}
