package instructions

import (
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations 
type invokespecial struct {Index16Instruction}
func (self *invokespecial) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    cp := frame.Method().Class().ConstantPool()
    cMethodRef := cp.GetConstant(self.index).(*class.ConstantMethodref)
    method := cMethodRef.SpecialMethod()
    newFrame := rtda.NewFrame(method)

    // pass args
    argCount := 1 + method.ArgCount()
    passArgs(frame.OperandStack(), newFrame.LocalVars(), argCount)

    thread.PushFrame(newFrame)
}
