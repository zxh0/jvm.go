package instructions

import (
    //"fmt"
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
    classLoader := fakeFields[0].(*rtc.ClassLoader)
    mainClassName := fakeFields[1].(string)
    args := fakeFields[2].([]string)

    classesToLoadAndInit := []string{
        "java/lang/String",
        "java/io/PrintStream",
        "jvmgo/SystemOut",
        mainClassName}

    for _, className := range classesToLoadAndInit {
        class := classLoader.LoadClass(className)
        if class.NotInitialized() {
            undoExec(thread, fakeRef)
            initClass(class, thread)
            return
        }
    }

    // create PrintStream

    // System.out
    stdout := classLoader.LoadClass("jvmgo/SystemOut").NewObj()
    sysClass := classLoader.LoadClass("java/lang/System")
    outField := sysClass.GetField("out", "Ljava/io/PrintStream;")
    outField.PutStaticValue(stdout)

    // exec main()
    mainClass := classLoader.LoadClass(mainClassName)
    mainMethod := mainClass.GetMainMethod()
    if mainMethod != nil {
        newFrame := rtda.NewFrame(mainMethod)
        thread.PushFrame(newFrame)
        // todo create args
        jArgs := rtc.NewRefArray(int32(len(args)))
        newFrame.LocalVars().SetRef(0, jArgs)
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
