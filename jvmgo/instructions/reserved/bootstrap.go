package reserved

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

// todo
var (
	_bootClasses     []string
	_classLoader     *heap.ClassLoader // todo
	_mainClassName   string
	_args            []string
	_mainThreadGroup *heap.Object
)

// Fake instruction to load and execute main class
type BOOTSTRAP struct{ base.NoOperandsInstruction }

func (self *BOOTSTRAP) Execute(frame *rtda.Frame) {
	thread := frame.Thread()

	if _classLoader == nil {
		_classLoader = heap.BootLoader()
		initVars(frame)
	}
	if bootClassesNotReady(thread) ||
		mainThreadNotReady(thread) ||
		jlSystemNotReady(thread) ||
		mainClassNotReady(thread) {

		return
	}

	execMain(thread)
}

func initVars(frame *rtda.Frame) {
	vars := frame.LocalVars()
	_mainClassName = vars.Get(0).(string)
	_args = vars.Get(1).([]string)
	_bootClasses = []string{
		"java/lang/Class",
		"java/lang/String",
		"java/lang/System",
		"java/lang/Thread",
		"java/lang/ThreadGroup",
		"java/io/PrintStream",
	}
}

func bootClassesNotReady(thread *rtda.Thread) bool {
	for _, className := range _bootClasses {
		class := _classLoader.LoadClass(className)
		if class.InitializationNotStarted() {
			undoExec(thread)
			thread.InitClass(class)
			return true
		}
	}
	return false
}

func mainClassNotReady(thread *rtda.Thread) bool {
	mainClass := _classLoader.LoadClass(_mainClassName)
	if mainClass.InitializationNotStarted() {
		undoExec(thread)
		thread.InitClass(mainClass)
		return true
	}
	return false
}

func mainThreadNotReady(thread *rtda.Thread) bool {
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
	if thread.JThread() == nil {
		undoExec(thread)
		threadClass := _classLoader.LoadClass("java/lang/Thread")
		mainThreadObj := threadClass.NewObjWithExtra(thread)
		mainThreadObj.SetFieldValue("priority", "I", int32(1))
		thread.HackSetJThread(mainThreadObj)

		initMethod := threadClass.GetConstructor("(Ljava/lang/ThreadGroup;Ljava/lang/String;)V")
		stack.PushRef(mainThreadObj)        // this
		stack.PushRef(_mainThreadGroup)     // group
		stack.PushRef(rtda.JString("main")) // name
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
	thread.CurrentFrame().RevertNextPC()
}

func execMain(thread *rtda.Thread) {
	thread.PopFrame()
	mainClass := _classLoader.LoadClass(_mainClassName)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		newFrame := thread.NewFrame(mainMethod)
		thread.PushFrame(newFrame)
		args := createArgs()
		newFrame.LocalVars().SetRef(0, args)
	} else {
		panic("no main method!") // todo
	}
}

func createArgs() *heap.Object {
	jArgs := make([]*heap.Object, len(_args))
	for i, arg := range _args {
		jArgs[i] = rtda.JString(arg)
	}

	return heap.NewRefArray2(_classLoader.JLStringClass(), jArgs)
}
