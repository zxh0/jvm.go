package reflect

import (
    . "jvmgo/any"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

func init() {
    _array(newArray, "newArray", "(Ljava/lang/Class;I)Ljava/lang/Object;")
}

func _array(method Any, name, desc string) {
    rtc.RegisterNativeMethod("java/lang/reflect/Array", name, desc, method)
}

// private static native Object newArray(Class<?> componentType, int length)
// throws NegativeArraySizeException;
// (Ljava/lang/Class;I)Ljava/lang/Object;
func newArray(frame *rtda.Frame) {
    vars := frame.LocalVars()
    componentType := vars.GetRef(0)
    length := vars.GetInt(1)
    if length < 0 {
        // todo
        panic("NegativeArraySizeException")
    }

    componentClass := componentType.Extra().(*rtc.Class)
    arrObj := componentClass.NewArray(uint(length))

    stack := frame.OperandStack()
    stack.PushRef(arrObj)
}
