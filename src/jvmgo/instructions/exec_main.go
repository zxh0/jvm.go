package instructions

import (
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// Fake instruction to load and execute main class
type exec_main struct {NoOperandsInstruction}
func (self *exec_main) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()
    fakeRef := stack.PopRef()
    fakeFields := fakeRef.Fields().([]Any)
    className := fakeFields[0].(string)
    classLoader := fakeFields[1].(*rtc.ClassLoader)

    mainClass := classLoader.LoadClass(className)
    if mainClass.IsInitialized() {
        // prepare to reexec this instruction
        frame.SetNextPC(thread.PC())
        stack.PushRef(fakeRef)
        // todo init class
        initClass(mainClass)
        return
    }

    // todo find main()
    if mainClass == nil {
        panic("!!!!!")
    } else {
        panic("gogogo::" + mainClass.SuperClassName())
    }
}

func initClass(class *rtc.Class) {
    //rtda.InitClass(class)
}
