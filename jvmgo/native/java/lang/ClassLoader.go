package lang

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_cl(findBootstrapClass, "findBootstrapClass", "(Ljava/lang/String;)Ljava/lang/Class;")
	_cl(findLoadedClass0, "findLoadedClass0", "(Ljava/lang/String;)Ljava/lang/Class;")
}

func _cl(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/lang/ClassLoader", name, desc, method)
}

// private native Class<?> findBootstrapClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func findBootstrapClass(frame *rtda.Frame) {
	vars := frame.LocalVars()
	//this := vars.GetThis()
	name := vars.GetRef(1)

	className := rtc.DotToSlash(rtda.GoString(name))
	class := rtc.BootLoader().LoadClass(className)

	// todo: init class?
	stack := frame.OperandStack()
	stack.PushRef(class.JClass())

	// todo
	if r := recover(); r != nil {
		frame.OperandStack().PushRef(nil)
	}
}

// private native final Class<?> findLoadedClass0(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func findLoadedClass0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	name := vars.GetRef(1)

	className := rtc.DotToSlash(rtda.GoString(name))

	if isAppClassLoader(this) {
		class := rtc.BootLoader().FindLoadedClass(className)
		if class != nil {
			frame.OperandStack().PushRef(class.JClass())
		} else {
			frame.OperandStack().PushRef(nil)
		}
		return
	}

	// todo
	frame.OperandStack().PushRef(nil)
}

// todo
func isAppClassLoader(loader *rtc.Obj) bool {
	return loader.Class().Name() == "sun/misc/Launcher$AppClassLoader"
}
