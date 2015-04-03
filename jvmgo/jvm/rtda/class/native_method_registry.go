package class

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
)

// todo
// cannot use package 'native' because of cycle import!
var registry = map[string]Any{}
var emptyNativeMethod Any

func SetEmptyNativeMethod(m Any) {
	emptyNativeMethod = m
}

func RegisterNativeMethod(className, methodName, methodDescriptor string, method Any) {
	key := className + "~" + methodName + "~" + methodDescriptor

	if _, ok := registry[key]; !ok {
		registry[key] = method
	} else {
		panic("native method:" + key + " has been registered !")
	}
}

func findNativeMethod(method *Method) Any {
	key := method.class.name + "~" + method.name + "~" + method.descriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if method.IsRegisterNatives() {
		return emptyNativeMethod
	}
	panic("native method not found: " + key)
}
