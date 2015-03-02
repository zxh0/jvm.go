package reflect

import (
	. "jvmgo/any"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

func init() {
	_nmai(invoke0, "invoke0", "(Ljava/lang/reflect/Method;Ljava/lang/Object;[Ljava/lang/Object;)Ljava/lang/Object;")
}

func _nmai(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/reflect/NativeMethodAccessorImpl", name, desc, method)
}

// private static native Object invoke0(Method method, Object o, Object[] os);
// (Ljava/lang/reflect/Method;Ljava/lang/Object;[Ljava/lang/Object;)Ljava/lang/Object;
func invoke0(frame *rtda.Frame) {
	stack := frame.OperandStack()
	if stack.IsEmpty() {
		frame.RevertNextPC()
		_invoke(frame)
		return
	}

	// catch return value
	if stack.IsEmpty() {
		stack.PushNull()
		return
	}
	retVal := stack.Pop()
	if retVal == nil {
		stack.PushNull()
		return
	}
	switch retVal.(type) {
		default: panic("asdfadsfadsfdsf")
	}
}

func _invoke(frame *rtda.Frame) {
	vars := frame.LocalVars()
	methodObj := vars.GetRef(0)
	obj := vars.GetRef(1)
	argArrObj := vars.GetRef(2)
	
	goMethod := getGoMethod(methodObj)
	args := convertArgs(obj, argArrObj, goMethod)

	stack := frame.OperandStack()
	stack.HackSetSlots(args)
	if stack.IsEmpty() {
		stack.HackSetSlots([]Any{nil}) // make room for return value
	}

	frame.Thread().InvokeMethod(goMethod)
}
