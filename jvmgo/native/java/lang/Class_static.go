package lang

import (
	"github.com/zxh0/jvm.go/jvmgo/jutil"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

func init() {
	_class(desiredAssertionStatus0, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z")
	_class(forName0, "forName0", "(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;")
	_class(getPrimitiveClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;")
}

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
// (Ljava/lang/Class;)Z
func desiredAssertionStatus0(frame *rtda.Frame) {
	// todo
	stack := frame.OperandStack()
	//stack.PopRef() // this
	stack.PushBoolean(false)
}

// private static native Class<?> forName0(String name, boolean initialize,
//                                         ClassLoader loader,
//                                         Class<?> caller)
//     throws ClassNotFoundException;
// (Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;
func forName0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	jName := vars.GetRef(0)
	initialize := vars.GetBoolean(1)
	//jLoader := vars.GetRef(2)

	goName := rtda.GoString(jName)
	goName = jutil.ReplaceAll(goName, ".", "/")
	goClass := frame.ClassLoader().LoadClass(goName)
	jClass := goClass.JClass()

	if initialize && goClass.InitializationNotStarted() {
		// undo forName0
		thread := frame.Thread()
		frame.SetNextPC(thread.PC())
		// init class
		thread.InitClass(goClass)
	} else {
		stack := frame.OperandStack()
		stack.PushRef(jClass)
	}
}

// static native Class<?> getPrimitiveClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func getPrimitiveClass(frame *rtda.Frame) {
	vars := frame.LocalVars()
	nameObj := vars.GetRef(0)

	name := rtda.GoString(nameObj)
	classLoader := frame.ClassLoader()
	class := classLoader.GetPrimitiveClass(name)
	classObj := class.JClass()

	stack := frame.OperandStack()
	stack.PushRef(classObj)
}
