package instructions

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    //rtc "jvmgo/rtda/class"
)

// Fake instruction to load and execute main class
type exec_main struct {NoOperandsInstruction}
func (self *exec_main) Execute(thread *rtda.Thread) {

    fakeRef := thread.CurrentFrame().OperandStack().PopRef()
    fakeFields := fakeRef.Fields().([]Any)
    className := fakeFields[0].(string)
    //classLoader := fakeFields[1]

    panic("!!!" + className)

    
    // bytes := ref.Fields().([]byte)
    // mainClassName := string(bytes)

    // cp := frame.Method().Class().ConstantPool()
    // cClass := cp.GetConstant(self.index).(rtc.ConstantClass)
    // class := cClass.Class()
    // if !class.IsInitialized() {
    //     // todo init class
    // }
}
