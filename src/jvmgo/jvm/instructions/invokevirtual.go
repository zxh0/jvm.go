package instructions

import (
    //"log"
    //. "jvmgo/any"
    //"jvmgo/native"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
)

// Invoke instance method; dispatch based on class
type invokevirtual struct {Index16Instruction}
func (self *invokevirtual) Execute(frame *rtda.Frame) {
    thread := frame.Thread()
    stack := frame.OperandStack()
    cp := frame.Method().Class().ConstantPool()
    cMethodRef := cp.GetConstant(self.index).(*rtc.ConstantMethodref)
    ref := stack.Top(cMethodRef.ArgCount())
    if ref == nil {
        panic("NPE")
    }

    method := cMethodRef.VirtualMethod(ref.(*rtc.Obj))
    if method.IsNative() {
        nativeMethod, ok := method.NativeMethod().(func(*rtda.Frame))
        if ok {
            nativeMethod(frame)
            return
        }
    }
    thread.InvokeMethod(method)
}
