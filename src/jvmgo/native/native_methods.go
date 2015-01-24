package native

var registry = map[string]NativeMethod{"x":nil}

func register(className, methodName, methodDescriptor string, method NativeMethod) {
    key := className + "~" + methodName + "~" + methodDescriptor
    registry[key] = method
}

func FindNativeMethod(className, methodName, methodDescriptor string) (NativeMethod) {
    key := className + "~" + methodName + "~" + methodDescriptor
    if method, found := registry[key]; found {
        return method
    } else {
        panic("native method not found: " + key)
    }
}
