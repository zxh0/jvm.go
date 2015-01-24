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
    method := cMethodRef.Method()
    newFrame := rtda.NewFrame(method)
    newFrame.Method()
    // todo
    panic("todo invokevirtual")
}
