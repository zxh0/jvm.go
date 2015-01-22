package instructions

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// Fake instruction to load and execute main class
type exec_main struct {NoOperandsInstruction}
func (self *exec_main) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    fakeRef := stack.PopRef()
    fakeFields := fakeRef.Fields().([]Any)
    className := fakeFields[0].(string)
    classLoader := fakeFields[1].(*rtc.ClassLoader)

    mainClass := classLoader.LoadClass(className)
    if mainClass.IsInitialized() {
        stack.PushRef(fakeRef) // undo stack pop
        // todo init class
        rtda.InitClass(mainClass)
        return
    }

    // todo find main()
    if mainClass == nil {
        panic("!!!!!")
    } else {
        panic("gogogo::" + mainClass.SuperClassName())
    }
    

    
    // bytes := ref.Fields().([]byte)
    // mainClassName := string(bytes)

    // cp := frame.Method().Class().ConstantPool()
    // cClass := cp.GetConstant(self.index).(rtc.ConstantClass)
    // class := cClass.Class()
    // if !class.IsInitialized() {
    //     // todo init class
    // }
}
