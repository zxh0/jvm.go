package lang

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_cl(getCaller, "getCaller", "(I)Ljava/lang/Class;")
}

func _cl(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/lang/ClassLoader", name, desc, method)
}

// private static native Class<?> getCaller(int i)
// (I)Ljava/lang/Class;
func getCaller(frame *rtda.Frame) {
	vars := frame.LocalVars()
	i := uint(vars.GetInt(0))

	callerClass := frame.Thread().TopFrameN(i).Method().Class().JClass()
	stack := frame.OperandStack()
	stack.PushRef(callerClass)
}
