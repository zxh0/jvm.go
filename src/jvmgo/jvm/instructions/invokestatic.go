package instructions

import (
	//"log"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

// Invoke a class (static) method
type invokestatic struct{ Index16Instruction }

func (self *invokestatic) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := frame
	currentMethod := currentFrame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	kMethodRef := cp.GetConstant(self.index).(*rtc.ConstantMethodref)
	method := kMethodRef.StaticMethod()

	// init class
	classOfMethod := method.Class()
	if classOfMethod.InitializationNotStarted() {
		if classOfMethod != currentClass || !currentMethod.IsClinit() {
			currentFrame.SetNextPC(thread.PC())
			thread.InitClass(classOfMethod)
			return
		}
	}

	thread.InvokeMethod(method)
}
