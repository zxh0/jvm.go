package reflect

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
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
	if goMethod.IsStatic() {
		if goMethod.Class().InitializationNotStarted() {
			frame.RevertNextPC()
			frame.Thread().InitClass(goMethod.Class())
			return
		}
	}

	if goMethod.IsAbstract() {
		goMethod = obj.Class().GetInstanceMethod(goMethod.Name(), goMethod.Descriptor())
	}

	args := convertArgs(obj, argArrObj, goMethod)
	// remember return type
	returnType := goMethod.ParsedDescriptor().ReturnType()
	vars.Set(0, returnType)

	stack := frame.OperandStack()
	if len(args) > 0 {
		stack.HackSetSlots(args)
	} else {
		// make room for return value
		stack.HackSetSlots([]Any{nil})
		stack.Pop()
	}

	frame.Thread().InvokeMethod(goMethod)
	if returnType.IsVoidType() {
		stack.PushNull()
	}
}

func _boxReturnValue(frame *rtda.Frame, returnType *rtc.FieldType) {
	switch returnType.Descriptor()[0] {
	case 'Z':
		_callValueOf(frame, "Z", "java/lang/Boolean")
	case 'B':
		_callValueOf(frame, "B", "java/lang/Byte")
	case 'C':
		_callValueOf(frame, "C", "java/lang/Character")
	case 'S':
		_callValueOf(frame, "S", "java/lang/Short")
	case 'I':
		_callValueOf(frame, "I", "java/lang/Integer")
	case 'J':
		_callValueOf(frame, "J", "java/lang/Long")
	case 'F':
		_callValueOf(frame, "F", "java/lang/Float")
	case 'D':
		_callValueOf(frame, "D", "java/lang/Double")
	default:
		panic("Not primitive type: " + returnType.Descriptor())
	}
}

func _callValueOf(frame *rtda.Frame, primitiveDescriptor, wrapperClassName string) {
	wrapperClass := frame.ClassLoader().LoadClass(wrapperClassName)
	d := "(" + primitiveDescriptor + ")L" + wrapperClassName + ";"
	valueOfMethod := wrapperClass.GetStaticMethod("valueOf", d)
	frame.Thread().InvokeMethod(valueOfMethod)

	// init wrapper class
	if wrapperClass.InitializationNotStarted() {
		frame.Thread().InitClass(wrapperClass)
	}
}
