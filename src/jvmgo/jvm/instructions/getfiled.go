package instructions

import (
	//"fmt"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

// Fetch field from object
type getfield struct{ Index16Instruction }

func (self *getfield) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		// todo NullPointerException
		panic("NPE")
	}

	cp := frame.Method().Class().ConstantPool()
	cFieldRef := cp.GetConstant(self.index).(*rtc.ConstantFieldref)
	field := cFieldRef.InstanceField()
	val := field.GetValue(ref)

	stack.Push(val)
}

// Get static field from class
type getstatic struct{ Index16Instruction }

func (self *getstatic) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()

	cp := currentClass.ConstantPool()
	cFieldRef := cp.GetConstant(self.index).(*rtc.ConstantFieldref)
	field := cFieldRef.StaticField()

	classOfField := field.Class()
	if classOfField.InitializationNotStarted() {
		if classOfField != currentClass || !currentMethod.IsClinit() {
			thread := frame.Thread()
			frame.SetNextPC(thread.PC()) // undo getstatic
			rtda.InitClass(field.Class(), thread)
			return
		}
	}

	val := field.GetStaticValue()
	frame.OperandStack().Push(val)
}
