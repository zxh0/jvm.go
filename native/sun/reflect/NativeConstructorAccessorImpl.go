package reflect

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_ncai(newInstance0, "newInstance0", "(Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;")
}

func _ncai(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("sun/reflect/NativeConstructorAccessorImpl", name, desc, method)
}

// private static native Object newInstance0(Constructor<?> c, Object[] os)
// throws InstantiationException, IllegalArgumentException, InvocationTargetException;
// (Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;
func newInstance0(frame *rtda.Frame) {
	constructorObj := frame.GetRefVar(0)
	argArrObj := frame.GetRefVar(1)

	goConstructor := getGoConstructor(constructorObj)
	goClass := goConstructor.Class
	if goClass.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread.InitClass(goClass)
		return
	}

	obj := goClass.NewObj()
	frame.PushRef(obj)

	// call <init>
	args := convertArgs(obj, argArrObj, goConstructor)
	frame.Thread.InvokeMethodWithShim(goConstructor, args)
}
