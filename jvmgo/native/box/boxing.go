package box

import (
	"github.com/zxh0/jvm.go/jvmgo/jutil"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func Unbox(obj *heap.Object, primitiveDescriptor string) interface{} {
	return obj.GetFieldValue("value", primitiveDescriptor)
}

// boxing primitive types
// primitive value must be on the top of operand stack
func Box(frame *rtda.Frame, primitiveDescriptor byte) {
	switch primitiveDescriptor {
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
		jutil.Panicf("Not primitive type: %v", primitiveDescriptor)
	}
}

func _callValueOf(frame *rtda.Frame, primitiveDescriptor, wrapperClassName string) {
	wrapperClass := heap.BootLoader().LoadClass(wrapperClassName)
	valueOfDescriptor := "(" + primitiveDescriptor + ")L" + wrapperClassName + ";"
	valueOfMethod := wrapperClass.GetStaticMethod("valueOf", valueOfDescriptor)
	frame.Thread().InvokeMethod(valueOfMethod)

	// init wrapper class
	if wrapperClass.InitializationNotStarted() {
		frame.Thread().InitClass(wrapperClass)
	}
}
