package reflect

import (
	"github.com/zxh0/jvm.go/jvmgo/native/box"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_nmai(invoke0, "invoke0", "(Ljava/lang/reflect/Method;Ljava/lang/Object;[Ljava/lang/Object;)Ljava/lang/Object;")
}

func _nmai(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("sun/reflect/NativeMethodAccessorImpl", name, desc, method)
}

// private static native Object invoke0(Method method, Object o, Object[] os);
// (Ljava/lang/reflect/Method;Ljava/lang/Object;[Ljava/lang/Object;)Ljava/lang/Object;
func invoke0(frame *rtda.Frame) {
	stack := frame.OperandStack()
	if stack.IsEmpty() {
		frame.RevertNextPC()
		_invokeMethod(frame)
	} else {
		returnType := frame.LocalVars().Get(0).(*heap.FieldType)
		if returnType.IsBaseType() && !returnType.IsVoidType() {
			primitiveDescriptor := returnType.Descriptor()[0]
			box.Box(frame, primitiveDescriptor) // todo
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
	if len(args) > 1 {
		stack.HackSetSlots(args)
	} else if len(args) > 0 {
		// make room for return value
		stack.HackSetSlots([]interface{}{args[0], nil})
		stack.PopRef()
	} else {
		// make room for return value
		stack.HackSetSlots([]interface{}{nil, nil})
		stack.PopRef()
		stack.PopRef()
	}

	frame.Thread().InvokeMethod(goMethod)
	if returnType.IsVoidType() {
		stack.PushNull()
	}
}
