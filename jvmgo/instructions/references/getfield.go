package references

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

// Fetch field from object
type GET_FIELD struct {
	base.Index16Instruction
	field *heap.Field
}

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	if self.field == nil {
		cp := frame.Method().Class().ConstantPool()
		kFieldRef := cp.GetConstant(self.Index).(*heap.ConstantFieldref)
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
