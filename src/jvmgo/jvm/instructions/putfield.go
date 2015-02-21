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
		frame.Thread().ThrowNPE()
		return
	}

	cp := frame.ConstantPool()
	kFieldRef := cp.GetConstant(self.index).(*rtc.ConstantFieldref)
	field := kFieldRef.InstanceField()

	field.PutValue(ref, val)
}

// Set static field in class
type putstatic struct{ Index16Instruction }

func (self *putstatic) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()

	cp := currentClass.ConstantPool()
	kFieldRef := cp.GetConstant(self.index).(*rtc.ConstantFieldref)
	field := kFieldRef.StaticField()

	classOfField := field.Class()
	if classOfField.InitializationNotStarted() {
		if classOfField != currentClass || !currentMethod.IsClinit() {
			thread := frame.Thread()
			frame.SetNextPC(thread.PC()) // undo putstatic
			thread.InitClass(classOfField)
			return
		}
	}

	val := frame.OperandStack().Pop()
	field.PutStaticValue(val)
}
