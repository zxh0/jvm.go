package reflect

import (
	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/native/box"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_nmai(invoke0, "invoke0", "(Ljava/lang/reflect/Method;Ljava/lang/Object;[Ljava/lang/Object;)Ljava/lang/Object;")
}

func _nmai(method native.Method, name, desc string) {
	native.Register("sun/reflect/NativeMethodAccessorImpl", name, desc, method)
}

// private static native Object invoke0(Method method, Object o, Object[] os);
// (Ljava/lang/reflect/Method;Ljava/lang/Object;[Ljava/lang/Object;)Ljava/lang/Object;
func invoke0(frame *rtda.Frame) {
	if frame.IsStackEmpty() {
		frame.RevertNextPC()
		_invokeMethod(frame)
	} else {
		returnType := frame.GetLocalVar(0).GetHack().(heap.TypeDescriptor)
		if returnType.IsBaseType() && !returnType.IsVoidType() {
			primitiveDescriptor := returnType[0]
			box.Box(frame, primitiveDescriptor) // todo
		}
	}
}

func _invokeMethod(frame *rtda.Frame) {
	methodObj := frame.GetRefVar(0)
	obj := frame.GetRefVar(1)
	argArrObj := frame.GetRefVar(2)

	goMethod := getGoMethod(methodObj)
	if goMethod.IsStatic() {
		if goMethod.Class.InitializationNotStarted() {
			frame.RevertNextPC()
			frame.Thread.InitClass(goMethod.Class)
			return
		}
	}

	if goMethod.IsAbstract() {
		goMethod = obj.Class.GetInstanceMethod(goMethod.Name, goMethod.Descriptor)
	}

	args := convertArgs(obj, argArrObj, goMethod)
	// remember return type
	returnType := goMethod.ReturnType
	frame.SetLocalVar(0, heap.NewHackSlot(returnType))

	if len(args) > 1 {
		frame.HackSetSlots(args)
	} else if len(args) > 0 {
		// make room for return value
		frame.HackSetSlots([]heap.Slot{args[0], heap.EmptySlot})
		frame.PopRef()
	} else {
		// make room for return value
		frame.HackSetSlots([]heap.Slot{heap.EmptySlot, heap.EmptySlot})
		frame.PopRef()
		frame.PopRef()
	}

	frame.Thread.InvokeMethod(goMethod)
	if returnType.IsVoidType() {
		frame.PushNull()
	}
}
