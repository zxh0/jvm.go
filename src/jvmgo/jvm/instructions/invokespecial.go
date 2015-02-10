package instructions

import (
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations 
type invokespecial struct {Index16Instruction}
func (self *invokespecial) Execute(frame *rtda.Frame) {
    thread := frame.Thread()
    cp := frame.Method().Class().ConstantPool()
    cMethodRef := cp.GetConstant(self.index).(*rtc.ConstantMethodref)
    method := cMethodRef.SpecialMethod()

    // if method.IsNative() {
    //     // exec native method
    //     nativeMethod, ok := method.NativeMethod().(func(*rtda.Frame))
    //     if ok {
    //         nativeMethod(frame)
    //         return
    //     }
    // }

    thread.InvokeMethod(method)
}
