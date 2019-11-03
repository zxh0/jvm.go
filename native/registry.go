package native

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// NativeMethod
type Method func(frame *rtda.Frame)

var registry = map[string]Method{}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}

func Register(className, methodName, methodDescriptor string, method Method) {
	key := className + "~" + methodName + "~" + methodDescriptor

	if _, ok := registry[key]; !ok {
		registry[key] = method
	} else {
		panic("native method:" + key + " has been registered !")
	}
}

func FindNativeMethod(method *heap.Method) Method {
	key := method.Class.Name + "~" + method.Name + "~" + method.Descriptor
	if nativeMethod, ok := registry[key]; ok {
		return nativeMethod
	}
	if method.IsRegisterNatives() || method.IsInitIDs() {
		return emptyNativeMethod
	}
	panic("native method not found: " + key)
}
