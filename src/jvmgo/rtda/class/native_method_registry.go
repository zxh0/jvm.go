package class

import . "jvmgo/any"

// todo 
// cannot use package 'native' because of cycle import!
var registry = map[string]Any{}

func RegisterNativeMethod(className, methodName, methodDescriptor string, method Any) {
    key := className + "~" + methodName + "~" + methodDescriptor
    registry[key] = method
}

func findNativeMethod(method *Method) (Any) {
    key := method.class.name + "~" + method.name + "~" + method.descriptor
    if method, found := registry[key]; found {
        return method
    } else {
        panic("native method not found: " + key)
    }
}
