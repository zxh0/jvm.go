package instructions

import (
    //"fmt"
    "jvmgo/jvm/rtda"
    rtc "jvmgo/jvm/rtda/class"
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
type bootstrap struct {NoOperandsInstruction}
func (self *bootstrap) Execute(frame *rtda.Frame) {
    thread := frame.Thread()
    stack := frame.OperandStack()

    if _classLoader == nil {
        initVars(stack)
        _classLoader.Init()
    }
    if bootClassesNotReady(thread) {
        return
    }
    if mainThreadNotReady(thread) {
        return
    }
    if jlSystemNotReady(thread) {
        return
    }

    // exec main()
    thread.PopFrame()
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
        _mainClassName}
}

func bootClassesNotReady(thread *rtda.Thread) (bool) {
    for _, className := range _basicClasses {
        class := _classLoader.LoadClass(className)
        if class.InitializationNotStarted() {
            undoExec(thread)
            rtda.InitClass(class, thread)
            return true
        }
    }
    return false
}

func mainThreadNotReady(thread *rtda.Thread) (bool) {
    stack := thread.CurrentFrame().OperandStack()
    if _mainThreadGroup == nil {
        undoExec(thread)
        threadGroupClass := _classLoader.LoadClass("java/lang/ThreadGroup")
        _mainThreadGroup = threadGroupClass.NewObj()
        initMethod := threadGroupClass.GetConstructor("()V")
        stack.PushRef(_mainThreadGroup) // this
        thread.InvokeMethod(initMethod)
        return true
    }
    if _mainThreadName == nil {
        undoExec(thread)
        _mainThreadName = rtda.NewJString("main", thread.CurrentFrame())
        return true
    }
    if thread.JThread() == nil {
        undoExec(thread)
        threadClass := _classLoader.LoadClass("java/lang/Thread")
        mainThread := threadClass.NewObj()
        threadClass.GetInstanceField("priority", "I").PutValue(mainThread, int32(1))
        thread.HackSetJThread(mainThread)

        initMethod := threadClass.GetConstructor("(Ljava/lang/ThreadGroup;Ljava/lang/String;)V")
        stack.PushRef(mainThread) // this
        stack.PushRef(_mainThreadGroup) // group
        stack.PushRef(_mainThreadName) // name
        thread.InvokeMethod(initMethod)
        return true
    }
    return false
}

func jlSystemNotReady(thread *rtda.Thread) bool {
    sysClass := _classLoader.LoadClass("java/lang/System")
    propsField := sysClass.GetStaticField("props", "Ljava/util/Properties;")
    props := propsField.GetStaticValue()
    if props == nil {
        undoExec(thread)
        initSys := sysClass.GetStaticMethod("initializeSystemClass", "()V")
        thread.InvokeMethod(initSys)
        return true
    }
    return false
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

    return rtc.NewRefArray2(_classLoader.StringClass(), jArgs)
}
