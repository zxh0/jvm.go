package heap

// todo
// cannot use package 'native' because of cycle import!
var registry = map[string]interface{}{}
var emptyNativeMethod interface{}

func SetEmptyNativeMethod(m interface{}) {
	emptyNativeMethod = m
}

func RegisterNativeMethod(className, methodName, methodDescriptor string, method interface{}) {
	key := className + "~" + methodName + "~" + methodDescriptor

	if _, ok := registry[key]; !ok {
		registry[key] = method
	} else {
		panic("native method:" + key + " has been registered !")
	}
}

func findNativeMethod(method *Method) interface{} {
	key := method.class.name + "~" + method.name + "~" + method.descriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if method.IsRegisterNatives() || method.IsInitIDs() {
		return emptyNativeMethod
	}
	panic("native method not found: " + key)
}
