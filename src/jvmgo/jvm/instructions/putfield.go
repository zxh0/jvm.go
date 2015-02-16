package instructions

import (
	//"fmt"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

// Set field in object
type putfield struct{ Index16Instruction }

func (self *putfield) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.Pop()
	ref := stack.PopRef()
	if ref == nil {
		// todo NullPointerException
		panic("NPE")
	}

	cp := frame.Method().Class().ConstantPool()
	cFieldRef := cp.GetConstant(self.index).(*rtc.ConstantFieldref)
	field := cFieldRef.InstanceField()

	field.PutValue(ref, val)
}

// Set static field in class
type putstatic struct{ Index16Instruction }

func (self *putstatic) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()

	cp := currentClass.ConstantPool()
	cFieldRef := cp.GetConstant(self.index).(*rtc.ConstantFieldref)
	field := cFieldRef.StaticField()

	classOfField := field.Class()
	if classOfField.InitializationNotStarted() {
		if classOfField != currentClass || !currentMethod.IsClinit() {
			thread := frame.Thread()
			frame.SetNextPC(thread.PC()) // undo putstatic
			rtda.InitClass(classOfField, thread)
			return
		}
	}

	val := frame.OperandStack().Pop()
	field.PutStaticValue(val)
}
