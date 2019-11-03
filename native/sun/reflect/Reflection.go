package reflect

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_reflection(getCallerClass, "getCallerClass", "()Ljava/lang/Class;")
	_reflection(getClassAccessFlags, "getClassAccessFlags", "(Ljava/lang/Class;)I")
}

func _reflection(method native.Method, name, desc string) {
	native.Register("sun/reflect/Reflection", name, desc, method)
}

// public static native Class<?> getCallerClass(int i);
// (I)Ljava/lang/Class;
func getCallerClass(frame *rtda.Frame) {
	// top0 is sun/reflect/Reflection
	// top1 is the caller of getCallerClass()
	// top2 is the caller of method
	callerFrame := frame.Thread.TopFrameN(2)
	callerClass := callerFrame.GetClass().JClass
	frame.PushRef(callerClass)
}

// public static native int getClassAccessFlags(Class<?> type);
// (Ljava/lang/Class;)I
func getClassAccessFlags(frame *rtda.Frame) {
	_type := frame.GetRefVar(0)
	goClass := _type.GetGoClass()
	frame.PushInt(int32(goClass.AccessFlags))
}
