package instructions

import (
    //"fmt"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

var _classLoader *rtc.ClassLoader
var mainClassName string
var args []string
var jArgs []*rtc.Obj

// Fake instruction to load and execute main class
type exec_main struct {NoOperandsInstruction}
func (self *exec_main) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()

    if _classLoader == nil {
        fakeRef := stack.PopRef()
        fakeFields := fakeRef.Fields().([]Any)
        _classLoader = fakeFields[0].(*rtc.ClassLoader)
        mainClassName = fakeFields[1].(string)
        args = fakeFields[2].([]string)
    }

    classesToLoadAndInit := []string{
        "java/lang/String",
        "java/io/PrintStream",
        "jvmgo/SystemOut",
        mainClassName}

    for _, className := range classesToLoadAndInit {
        class := _classLoader.LoadClass(className)
        if class.NotInitialized() {
            undoExec(thread)
            initClass(class, thread)
            return
        }
    }

    // create args
    if len(args) > 0 {
        if jArgs == nil {
            jArgs = make([]*rtc.Obj, 0, len(args))
        } else {
            jArgs = jArgs[:len(jArgs) + 1]
            jArgs[len(jArgs) - 1] = stack.PopRef()
        }
        for len(jArgs) < len(args) {
            undoExec(thread)
            newJString(args[len(jArgs)], thread)
            return
        }
    }

    // create PrintStream

    // System.out
    stdout := _classLoader.LoadClass("jvmgo/SystemOut").NewObj()
    sysClass := _classLoader.LoadClass("java/lang/System")
    outField := sysClass.GetField("out", "Ljava/io/PrintStream;")
    outField.PutStaticValue(stdout)

    // exec main()
    mainClass := _classLoader.LoadClass(mainClassName)
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
func undoExec(thread *rtda.Thread) {
    thread.CurrentFrame().SetNextPC(thread.PC())
}
