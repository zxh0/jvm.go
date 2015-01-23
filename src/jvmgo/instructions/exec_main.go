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

    // load and init java.lang.String
    stringClass := classLoader.LoadClass("java/lang/String")
    if stringClass.IsInitialized() {
        undoExec(thread, fakeRef)
        initClass(stringClass, thread)
        return
    }

    // load and init main class
    mainClass := classLoader.LoadClass(className)
    if !mainClass.IsInitialized() {
        undoExec(thread, fakeRef)
        initClass(mainClass, thread)
        return
    }

    // exec main()
    mainMethod := mainClass.GetMainMethod()
    if mainMethod != nil {
        newFrame := rtda.NewFrame(mainMethod)
        thread.PushFrame(newFrame)
        // todo create args
        panic("here!!!")
    } else {
        panic("no main method!")
    }
}

// prepare to reexec this instruction
func undoExec(thread *rtda.Thread, fakeRef *rtc.Obj) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()
    frame.SetNextPC(thread.PC())
    stack.PushRef(fakeRef)
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
