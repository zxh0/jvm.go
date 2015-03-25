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
		cp := frame.Method().Class().ConstantPool()
		kFieldRef := cp.GetConstant(self.index).(*rtc.ConstantFieldref)
		self.field = kFieldRef.InstanceField()
	}

	stack := frame.OperandStack()
	val := stack.PopField(self.field.IsLongOrDouble)
	ref := stack.PopRef()
	if ref == nil {
		frame.Thread().ThrowNPE()
		return
	}

	self.field.PutValue(ref, val)
}
