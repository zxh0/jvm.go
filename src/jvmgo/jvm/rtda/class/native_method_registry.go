package class

import . "jvmgo/any"

// todo
// cannot use package 'native' because of cycle import!
var registry = map[string]Any{}
var registerNatives Any

func SetRegisterNatives(_registerNatives Any) {
	registerNatives = _registerNatives
}

func RegisterNativeMethod(className, methodName, methodDescriptor string, method Any) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

func findNativeMethod(method *Method) Any {
	if method.IsRegisterNatives() {
		return registerNatives
	}

	key := method.class.name + "~" + method.name + "~" + method.descriptor
	if method, ok := registry[key]; ok {
		return method
	} else {
		panic("native method not found: " + key)
	}
}
