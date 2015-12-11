package security

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_ac(doPrivileged, "doPrivileged", "(Ljava/security/PrivilegedAction;)Ljava/lang/Object;")
	_ac(doPrivileged2, "doPrivileged", "(Ljava/security/PrivilegedAction;Ljava/security/AccessControlContext;)Ljava/lang/Object;")
	_ac(doPrivileged3, "doPrivileged", "(Ljava/security/PrivilegedExceptionAction;)Ljava/lang/Object;")
	_ac(doPrivileged4, "doPrivileged", "(Ljava/security/PrivilegedExceptionAction;Ljava/security/AccessControlContext;)Ljava/lang/Object;")
	_ac(getStackAccessControlContext, "getStackAccessControlContext", "()Ljava/security/AccessControlContext;")
}

func _ac(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/security/AccessController", name, desc, method)
}

// @CallerSensitive
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
// public static native <T> T doPrivileged(PrivilegedAction<T> action,
//                                         AccessControlContext context);
// (Ljava/security/PrivilegedAction;Ljava/security/AccessControlContext;)Ljava/lang/Object;
func doPrivileged2(frame *rtda.Frame) {
	// todo
	doPrivileged(frame)
}

// @CallerSensitive
// public static native <T> T
//     doPrivileged(PrivilegedExceptionAction<T> action)
//     throws PrivilegedActionException;
// (Ljava/security/PrivilegedExceptionAction;)Ljava/lang/Object;
func doPrivileged3(frame *rtda.Frame) {
	// todo
	doPrivileged(frame)
}

// @CallerSensitive
// public static native <T> T
//     doPrivileged(PrivilegedExceptionAction<T> action,
//                  AccessControlContext context)
//     throws PrivilegedActionException;
// (Ljava/security/PrivilegedExceptionAction;Ljava/security/AccessControlContext;)Ljava/lang/Object;
func doPrivileged4(frame *rtda.Frame) {
	// todo
	doPrivileged(frame)
}

// private static native AccessControlContext getStackAccessControlContext();
// ()Ljava/security/AccessControlContext;
func getStackAccessControlContext(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushRef(nil)
}
