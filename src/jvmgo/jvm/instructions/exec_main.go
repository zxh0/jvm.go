package instructions

import (
    //"fmt"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// todo
var (
    _basicClasses       []string
    _classLoader        *rtc.ClassLoader
    _mainClassName      string
    _args               []string
    _mainThreadGroup    *rtc.Obj
    _mainThreadName     *rtc.Obj
)

// Fake instruction to load and execute main class
type exec_main struct {NoOperandsInstruction}
func (self *exec_main) Execute(frame *rtda.Frame) {
    thread := frame.Thread()
    stack := frame.OperandStack()

    if _classLoader == nil {
        initVars(stack)
        _classLoader.Init()
    }
    if !isBasicClassesReady(thread) {
        return
    }
    if !isMainThreadReady(thread) {
        return
    }
    
    // todo create PrintStream

    // System.out
    sysClass := _classLoader.LoadClass("java/lang/System")
    propsField := sysClass.GetField("props", "Ljava/util/Properties;")
    props := propsField.GetStaticValue()
    if props == nil {
        // undoExec(thread)
        // initSys := sysClass.GetStaticMethod("initializeSystemClass", "()V")
        // thread.InvokeMethod(initSys)
        // return
    }

    outField := sysClass.GetField("out", "Ljava/io/PrintStream;")
    stdout := _classLoader.LoadClass("jvmgo/SystemOut").NewObj()
    outField.PutStaticValue(stdout)

    // exec main()
    mainClass := _classLoader.LoadClass(_mainClassName)
    mainMethod := mainClass.GetMainMethod()
    if mainMethod != nil {
        newFrame := thread.NewFrame(mainMethod)
        thread.PushFrame(newFrame)
        args := createArgs(newFrame)
        newFrame.LocalVars().SetRef(0, args)
    } else {
        panic("no main method!") // todo
    }
}

func initVars(stack *rtda.OperandStack) {
    _classLoader = stack.Pop().(*rtc.ClassLoader)
    _mainClassName = stack.Pop().(string)
    _args = stack.Pop().([]string)
    _basicClasses = []string{
        "java/lang/Class",
        "java/lang/String",
        "java/lang/System",
        "java/lang/Thread",
        "java/lang/ThreadGroup",
        "java/io/PrintStream",
        "jvmgo/SystemOut",
        _mainClassName}
}

func isBasicClassesReady(thread *rtda.Thread) (bool) {
    for _, className := range _basicClasses {
        class := _classLoader.LoadClass(className)
        if class.InitializationNotStarted() {
            undoExec(thread)
            initClass(class, thread)
            return false
        }
    }
    return true
}

func isMainThreadReady(thread *rtda.Thread) (bool) {
    stack := thread.CurrentFrame().OperandStack()
    if _mainThreadGroup == nil {
        undoExec(thread)
        threadGroupClass := _classLoader.LoadClass("java/lang/ThreadGroup")
        _mainThreadGroup = threadGroupClass.NewObj()
        initMethod := threadGroupClass.GetMethod("<init>", "()V")
        stack.PushRef(_mainThreadGroup) // this
        thread.InvokeMethod(initMethod)
        return false
    }
    if _mainThreadName == nil {
        undoExec(thread)
        _mainThreadName = rtda.NewJString("main", thread.CurrentFrame())
        return false
    }
    if thread.JThread() == nil {
        undoExec(thread)
        threadClass := _classLoader.LoadClass("java/lang/Thread")
        mainThread := threadClass.NewObj()
        threadClass.GetField("priority", "I").PutValue(mainThread, int32(1))
        thread.SetJThread(mainThread)

        initMethod := threadClass.GetMethod("<init>", "(Ljava/lang/ThreadGroup;Ljava/lang/String;)V")
        stack.PushRef(mainThread) // this
        stack.PushRef(_mainThreadGroup) // group
        stack.PushRef(_mainThreadName) // name
        thread.InvokeMethod(initMethod)
        return false
    }
    return true
}

// prepare to reexec this instruction
func undoExec(thread *rtda.Thread) {
    thread.CurrentFrame().SetNextPC(thread.PC())
}

func createArgs(frame *rtda.Frame) (*rtc.Obj) {
    jArgs := make([]*rtc.Obj, len(_args))
    for i, arg := range _args {
        jArgs[i] = rtda.NewJString(arg, frame)
    }

    return rtc.NewRefArray2(_classLoader.StringClass(), jArgs, _classLoader)
}
