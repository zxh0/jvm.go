package instructions

import (
    "jvmgo/rtda"
    //rtc "jvmgo/rtda/class"
)

// Fake instruction to trigger JVM startup
type main struct {NoOperandsInstruction}
func (self *main) Execute(thread *rtda.Thread) {
    // ref := thread.CurrentFrame().OperandStack().PopRef()
    // bytes := ref.Fields().([]byte)
    // mainClassName := string(bytes)

    // cp := frame.Method().Class().ConstantPool()
    // cClass := cp.GetConstant(self.index).(rtc.ConstantClass)
    // class := cClass.Class()
    // if !class.IsInitialized() {
    //     // todo init class
    // }
}
