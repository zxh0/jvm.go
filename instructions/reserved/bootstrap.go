package reserved

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Fake instruction to load and execute main class
type Bootstrap struct {
	base.NoOperandsInstruction

	bootLoader      *heap.ClassLoader // todo
	mainClassName   string
	args            []string
	bootClassNames  []string
	mainThreadGroup *heap.Object
}

func (instr *Bootstrap) Execute(frame *rtda.Frame) {
	if instr.bootLoader == nil {
		instr.init(frame)
	}

	thread := frame.Thread
	if instr.bootClassesNotReady(thread) ||
		instr.mainThreadNotReady(thread) ||
		instr.jlSystemNotReady(thread) ||
		instr.mainClassNotReady(thread) {

		return
	}

	instr.execMain(thread)
}

func (instr *Bootstrap) init(frame *rtda.Frame) {
	instr.bootLoader = frame.GetBootLoader()
	instr.mainClassName = frame.GetLocalVar(0).GetHack().(string)
	instr.args = frame.GetLocalVar(1).GetHack().([]string)
	instr.bootClassNames = []string{
		"java/lang/Class",
		"java/lang/String",
		"java/lang/System",
		"java/lang/Thread",
		"java/lang/ThreadGroup",
		"java/io/PrintStream",
	}
}

func (instr *Bootstrap) bootClassesNotReady(thread *rtda.Thread) bool {
	for _, className := range instr.bootClassNames {
		class := instr.bootLoader.LoadClass(className)
		if class.InitializationNotStarted() {
			undoExec(thread)
			thread.InitClass(class)
			return true
		}
	}
	return false
}

func (instr *Bootstrap) mainClassNotReady(thread *rtda.Thread) bool {
	mainClass := instr.bootLoader.LoadClass(instr.mainClassName)
	if mainClass.InitializationNotStarted() {
		undoExec(thread)
		thread.InitClass(mainClass)
		return true
	}
	return false
}

func (instr *Bootstrap) mainThreadNotReady(thread *rtda.Thread) bool {
	frame := thread.CurrentFrame()
	if instr.mainThreadGroup == nil {
		undoExec(thread)
		threadGroupClass := instr.bootLoader.LoadClass("java/lang/ThreadGroup")
		instr.mainThreadGroup = threadGroupClass.NewObj()
		initMethod := threadGroupClass.GetConstructor("()V")
		frame.PushRef(instr.mainThreadGroup) // this
		thread.InvokeMethod(initMethod)
		return true
	}
	if thread.JThread() == nil {
		undoExec(thread)
		threadClass := instr.bootLoader.LoadClass("java/lang/Thread")
		mainThreadObj := threadClass.NewObjWithExtra(thread)
		mainThreadObj.SetFieldValue("priority", "I", heap.NewIntSlot(1))
		thread.HackSetJThread(mainThreadObj)

		initMethod := threadClass.GetConstructor("(Ljava/lang/ThreadGroup;Ljava/lang/String;)V")
		frame.PushRef(mainThreadObj)                          // this
		frame.PushRef(instr.mainThreadGroup)                  // group
		frame.PushRef(frame.GetRuntime().JSFromGoStr("main")) // name
		thread.InvokeMethod(initMethod)
		return true
	}
	return false
}

func (instr *Bootstrap) jlSystemNotReady(thread *rtda.Thread) bool {
	sysClass := instr.bootLoader.LoadClass("java/lang/System")
	propsField := sysClass.GetStaticField("props", "Ljava/util/Properties;")
	props := propsField.GetStaticValue().Ref
	if props == nil {
		undoExec(thread)
		initSys := sysClass.GetStaticMethod("initializeSystemClass", "()V")
		thread.InvokeMethod(initSys)
		return true
	}
	return false
}

func (instr *Bootstrap) execMain(thread *rtda.Thread) {
	thread.PopFrame()
	mainClass := instr.bootLoader.LoadClass(instr.mainClassName)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		newFrame := thread.NewFrame(mainMethod)
		thread.PushFrame(newFrame)
		args := instr.createArgs(thread.Runtime)
		newFrame.SetRefVar(0, args)
	} else {
		panic("no main method!") // todo
	}
}

func (instr *Bootstrap) createArgs(rt *heap.Runtime) *heap.Object {
	jArgs := make([]*heap.Object, len(instr.args))
	for i, arg := range instr.args {
		jArgs[i] = rt.JSFromGoStr(arg)
	}

	return rt.NewStringArray(jArgs)
}

// prepare to reexec this instruction
func undoExec(thread *rtda.Thread) {
	thread.CurrentFrame().RevertNextPC()
}
