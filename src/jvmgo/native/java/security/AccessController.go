package security

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_ac(doPrivileged, "doPrivileged", "(Ljava/security/PrivilegedAction;)Ljava/lang/Object;")
	_ac(doPrivileged2, "doPrivileged", "(Ljava/security/PrivilegedExceptionAction;)Ljava/lang/Object;")
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

	stack := frame.OperandStack()
	stack.PushRef(action)

	method := action.Class().GetInstanceMethod("run", "()Ljava/lang/Object;") // todo
	frame.Thread().InvokeMethod(method)
}

// @CallerSensitive
// public static native <T> T
//     doPrivileged(PrivilegedExceptionAction<T> action)
//     throws PrivilegedActionException;
// (Ljava/security/PrivilegedExceptionAction;)Ljava/lang/Object;
func doPrivileged2(frame *rtda.Frame) {
	// todo
	doPrivileged(frame)
}

// private static native AccessControlContext getStackAccessControlContext();
// ()Ljava/security/AccessControlContext;
func getStackAccessControlContext(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushRef(nil)
}
