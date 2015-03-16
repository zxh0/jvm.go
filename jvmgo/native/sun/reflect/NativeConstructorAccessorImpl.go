package reflect

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_ncai(newInstance0, "newInstance0", "(Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;")
}

func _ncai(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/reflect/NativeConstructorAccessorImpl", name, desc, method)
}

// private static native Object newInstance0(Constructor<?> c, Object[] os)
// throws InstantiationException, IllegalArgumentException, InvocationTargetException;
// (Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;
func newInstance0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	constructorObj := vars.GetRef(0)
	argArrObj := vars.GetRef(1)

	goConstructor := getGoConstructor(constructorObj)
	goClass := goConstructor.Class()
	if goClass.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread().InitClass(goClass)
		return
	}

	obj := goClass.NewObj()
	stack := frame.OperandStack()
	stack.PushRef(obj)

	// call <init>
	args := convertArgs(obj, argArrObj, goConstructor)
	frame.Thread().InvokeMethodWithShim(goConstructor, args)
}
