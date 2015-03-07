package lang

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_cl(findLoadedClass0, "findLoadedClass0", "(Ljava/lang/String;)Ljava/lang/Class;")
}

func _cl(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/lang/ClassLoader", name, desc, method)
}

// private native final Class<?> findLoadedClass0(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func findLoadedClass0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	//this := vars.GetThis()
	name := vars.GetRef(1)

	// todo
	className := rtc.DotToSlash(rtda.GoString(name))
	class := frame.ClassLoader().LoadClass(className)
	if class != nil {
		frame.OperandStack().PushRef(class.JClass())
	} else {
		frame.OperandStack().PushRef(nil)
	}
}
