package instructions

import (
    "log"
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
    if !mainClass.IsInitialized() {
        // prepare to reexec this instruction
        frame.SetNextPC(thread.PC())
        stack.PushRef(fakeRef)
        // init class
        initClass(mainClass, thread)
    } else {
        // exec main()
        mainMethod := mainClass.GetMainMethod()
        if mainMethod != nil {
            panic("here!!!")
            newFrame := rtda.NewFrame(mainMethod)
            thread.PushFrame(newFrame)
        } else {
            panic("no main method!")
        }
    }
}

func initClass(class *rtc.Class, thread *rtda.Thread) {
    uninitedClass := rtc.GetUpmostUninitializedClassOrInterface(class)
    if uninitedClass != nil {
        log.Printf("init class: %v", uninitedClass.Name())
        clinit := uninitedClass.GetClinitMethod()
        if clinit != nil {
            // exec <clinit>
            newFrame := rtda.NewFrame(clinit)
            newFrame.SetOnPopAction(func() {
                uninitedClass.MarkInitialized()
            })
            thread.PushFrame(newFrame)
        } else {
            // no <clinit> method
            uninitedClass.MarkInitialized()
        }
    }
}
