package instructions

import (
    //"fmt"
    . "jvmgo/any"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// todo
var (
    _basicClasses []string
    _classLoader *rtc.ClassLoader
    _mainClassName string
    _args []string
    _jArgs []*rtc.Obj
)

// Fake instruction to load and execute main class
type exec_main struct {NoOperandsInstruction}
func (self *exec_main) Execute(frame *rtda.Frame) {
    thread := frame.Thread()
    stack := frame.OperandStack()

    if _classLoader == nil {
        initVars(stack.PopRef())
        _classLoader.Init()
    }
    if !isBasicClassesReady(thread) {
        return
    }
    if !isJArgsReady(thread) {
        return
    }
    
    // todo create PrintStream

    // System.out
    stdout := _classLoader.LoadClass("jvmgo/SystemOut").NewObj()
    sysClass := _classLoader.LoadClass("java/lang/System")
    outField := sysClass.GetField("out", "Ljava/io/PrintStream;")
    outField.PutStaticValue(stdout)

    // exec main()
    mainClass := _classLoader.LoadClass(_mainClassName)
    mainMethod := mainClass.GetMainMethod()
    if mainMethod != nil {
        newFrame := thread.NewFrame(mainMethod)
        thread.PushFrame(newFrame)
        args := rtc.NewRefArrayOfElements(_jArgs)
        newFrame.LocalVars().SetRef(0, args)
    } else {
        panic("no main method!") // todo
    }
}

func initVars(fakeRef *rtc.Obj) {
    fakeFields := fakeRef.Fields().([]Any)
    _classLoader = fakeFields[0].(*rtc.ClassLoader)
    _mainClassName = fakeFields[1].(string)
    _args = fakeFields[2].([]string)
    _basicClasses = []string{
        "java/lang/Class",
        "java/lang/String",
        "java/io/PrintStream",
        "jvmgo/SystemOut",
        _mainClassName}
}

func isBasicClassesReady(thread *rtda.Thread) (bool) {
    for _, className := range _basicClasses {
        class := _classLoader.LoadClass(className)
        if class.NotInitialized() {
            undoExec(thread)
            initClass(class, thread)
            return false
        }
    }
    return true
}

func isJArgsReady(thread *rtda.Thread) (bool) {
    if len(_args) > 0 {
        if _jArgs == nil {
            _jArgs = make([]*rtc.Obj, 0, len(_args))
        } else {
            jStr := thread.CurrentFrame().OperandStack().PopRef()
            _jArgs = _jArgs[:len(_jArgs) + 1]
            _jArgs[len(_jArgs) - 1] = jStr
        }
        for len(_jArgs) < len(_args) {
            undoExec(thread)
            newJString(_args[len(_jArgs)], thread)
            return false
        }
    }
    return true
}

// prepare to reexec this instruction
func undoExec(thread *rtda.Thread) {
    thread.CurrentFrame().SetNextPC(thread.PC())
}
