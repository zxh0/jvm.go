package reflect

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
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

	goConstructor := getExtra(constructorObj)
	goClass := goConstructor.Class()
	obj := goClass.NewObj()
	stack := frame.OperandStack()
	stack.PushRef(obj)

	// call <init>
	args := convertArgs(obj, argArrObj, goConstructor)
	frame.Thread().InvokeMethodWithShim(goConstructor, args)
}

func getExtra(constructorObj *rtc.Obj) *rtc.Method {
	extra := constructorObj.Extra()
	if extra != nil {
		return extra.(*rtc.Method)
	}

	root := constructorObj.GetFieldValue("root", "Ljava/lang/reflect/Constructor;").(*rtc.Obj)
	return root.Extra().(*rtc.Method)
}
