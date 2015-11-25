package references

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

// Fetch field from object
type GET_FIELD struct {
	base.Index16Instruction
	field *rtc.Field
}

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	if self.field == nil {
		cp := frame.Method().Class().ConstantPool()
		kFieldRef := cp.GetConstant(self.Index).(*rtc.ConstantFieldref)
		self.field = kFieldRef.InstanceField()
	}

	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		frame.Thread().ThrowNPE()
		return
	}

	val := self.field.GetValue(ref)
	stack.PushField(val, self.field.IsLongOrDouble)
}
