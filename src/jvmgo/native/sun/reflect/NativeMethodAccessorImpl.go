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
		_invokeMethod(frame)
	} else {
		returnType := frame.LocalVars().Get(0).(*rtc.FieldType)
		if returnType.IsBaseType() && !returnType.IsVoidType() {
			_boxReturnValue(frame, returnType)
		}
	}
}

func _invokeMethod(frame *rtda.Frame) {
	vars := frame.LocalVars()
	methodObj := vars.GetRef(0)
	obj := vars.GetRef(1)
	argArrObj := vars.GetRef(2)
	
	goMethod := getGoMethod(methodObj)
	args := convertArgs(obj, argArrObj, goMethod)
	// remember boolean return type
	returnType := goMethod.ParsedDescriptor().ReturnType()
	vars.Set(0, returnType)

	stack := frame.OperandStack()
	if len(args) > 0 {
		stack.HackSetSlots(args)
	} else {
		// make room for return value
		stack.HackSetSlots([]Any{nil})
	}

	frame.Thread().InvokeMethod(goMethod)
	if returnType.IsVoidType() {
		stack.PushNull()
	}
}

func _boxReturnValue(frame *rtda.Frame, returnType *rtc.FieldType) {
	stack := frame.OperandStack()

	switch returnType.Descriptor()[0] {
		case 'B': _boxBoolean(stack.PopBoolean(), frame)
		case 'I': _boxInt(stack.PopInt(), frame)
	}
}

func _boxBoolean(val bool, frame *rtda.Frame) {
	// todo init boolean class?
	booleanClass := frame.ClassLoader().LoadClass("java/lang/Boolean")
	if booleanClass.InitializationNotStarted() {
		panic("todo: init java/lang/Boolean")
	}

	if val {
		boxed := booleanClass.GetStaticValue("TRUE", "Ljava/lang/Boolean;").(*rtc.Obj)
		frame.OperandStack().PushRef(boxed)
	} else {
		boxed := booleanClass.GetStaticValue("FALSE", "Ljava/lang/Boolean;").(*rtc.Obj)
		frame.OperandStack().PushRef(boxed)
	}
}

func _boxInt(val int32, frame *rtda.Frame) {
	// todo
}
