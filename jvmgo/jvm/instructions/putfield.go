package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

// Set field in object
type putfield struct {
	Index16Instruction
	field *rtc.Field
}

func (self *putfield) Execute(frame *rtda.Frame) {
	if self.field == nil {
		cp := frame.ConstantPool()
		kFieldRef := cp.GetConstant(self.index).(*rtc.ConstantFieldref)
		self.field = kFieldRef.InstanceField()
	}

	stack := frame.OperandStack()
	val := stack.Pop()
	ref := stack.PopRef()
	if ref == nil {
		frame.Thread().ThrowNPE()
		return
	}

	self.field.PutValue(ref, val)
}

// Set static field in class
type putstatic struct {
	Index16Instruction
	field *rtc.Field
}

func (self *putstatic) Execute(frame *rtda.Frame) {
	if self.field == nil {
		cp := frame.Method().Class().ConstantPool()
		kFieldRef := cp.GetConstant(self.index).(*rtc.ConstantFieldref)
		self.field = kFieldRef.StaticField()
	}

	class := self.field.Class()
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread().InitClass(class)
		return
	}

	val := frame.OperandStack().Pop()
	self.field.PutStaticValue(val)
}
