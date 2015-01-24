package instructions

import (
    //"log"
    //. "jvmgo/any"
    //"jvmgo/native"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// Invoke instance method; dispatch based on class
type invokevirtual struct {Index16Instruction}
func (self *invokevirtual) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    cp := frame.Method().Class().ConstantPool()
    cMethodRef := cp.GetConstant(self.index).(*rtc.ConstantMethodref)
    method := cMethodRef.VirtualMethod()
    newFrame := rtda.NewFrame(method)

    // pass args
    argCount := 1 + method.ArgCount()
    passArgs(frame.OperandStack(), newFrame.LocalVars(), argCount)

    thread.PushFrame(newFrame)
    // todo
    //panic("todo invokevirtual")
}
