package reflect

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_reflection(getCallerClass, "getCallerClass", "()Ljava/lang/Class;")
	_reflection(getClassAccessFlags, "getClassAccessFlags", "(Ljava/lang/Class;)I")
}

func _reflection(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/reflect/Reflection", name, desc, method)
}

// public static native Class<?> getCallerClass(int i);
// (I)Ljava/lang/Class;
func getCallerClass(frame *rtda.Frame) {
	// top0 is sun/reflect/Reflection
	// top1 is the caller of getCallerClass()
	// top2 is the caller of method
	callerFrame := frame.Thread().TopFrameN(2)
	callerClass := callerFrame.Method().Class().JClass()
	frame.OperandStack().PushRef(callerClass)
}

// public static native int getClassAccessFlags(Class<?> type);
// (Ljava/lang/Class;)I
func getClassAccessFlags(frame *rtda.Frame) {
	vars := frame.LocalVars()
	_type := vars.GetRef(0)

	goClass := _type.Extra().(*rtc.Class)
	flags := goClass.GetAccessFlags()

	stack := frame.OperandStack()
	stack.PushInt(int32(flags))
}
