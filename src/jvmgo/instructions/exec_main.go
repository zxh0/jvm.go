package instructions

import (
    "jvmgo/rtda"
    //rtc "jvmgo/rtda/class"
)

// Fake instruction to load and execute main class
type exec_main struct {NoOperandsInstruction}
func (self *exec_main) Execute(thread *rtda.Thread) {

    panic("!!!")

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
