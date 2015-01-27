package instructions

import (
    //"fmt"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

var (
    _classLoader *rtc.ClassLoader
    _mainClassName string
    _args []string
    _jArgs []*rtc.Obj
)

// Fake instruction to load and execute main class
type exec_main struct {NoOperandsInstruction}
func (self *exec_main) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()

    if _classLoader == nil {
        initVars(stack.PopRef())
    }

    classesToLoadAndInit := []string{
        "java/lang/String",
        "java/io/PrintStream",
        "jvmgo/SystemOut",
        _mainClassName}

    for _, className := range classesToLoadAndInit {
        class := _classLoader.LoadClass(className)
        if class.NotInitialized() {
            undoExec(thread)
            initClass(class, thread)
            return
        }
    }

    // create args
    if len(_args) > 0 {
        if _jArgs == nil {
            _jArgs = make([]*rtc.Obj, 0, len(_args))
        } else {
            _jArgs = _jArgs[:len(_jArgs) + 1]
            _jArgs[len(_jArgs) - 1] = stack.PopRef()
        }
        for len(_jArgs) < len(_args) {
            undoExec(thread)
            newJString(_args[len(_jArgs)], thread)
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
    mainClass := _classLoader.LoadClass(_mainClassName)
    mainMethod := mainClass.GetMainMethod()
    if mainMethod != nil {
        newFrame := rtda.NewFrame(mainMethod)
        thread.PushFrame(newFrame)
        // todo create args
        _jArgs := rtc.NewRefArray(int32(len(_args)))
        newFrame.LocalVars().SetRef(0, _jArgs)
    } else {
        panic("no main method!")
    }
}

func initVars(fakeRef *rtc.Obj) {
    fakeFields := fakeRef.Fields().([]Any)
    _classLoader = fakeFields[0].(*rtc.ClassLoader)
    _mainClassName = fakeFields[1].(string)
    _args = fakeFields[2].([]string)
}

// prepare to reexec this instruction
func undoExec(thread *rtda.Thread) {
    thread.CurrentFrame().SetNextPC(thread.PC())
}
