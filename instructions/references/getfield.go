package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Fetch field from object
type GET_FIELD struct {
	base.Index16Instruction
	field *heap.Field
}

func (instr *GET_FIELD) Execute(frame *rtda.Frame) {
	if instr.field == nil {
		cp := frame.Method().Class().ConstantPool()
		kFieldRef := cp.GetConstant(instr.Index).(*heap.ConstantFieldref)
		instr.field = kFieldRef.InstanceField()
	}

	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		frame.Thread().ThrowNPE()
		return
	}

	val := instr.field.GetValue(ref)
	stack.PushField(val, instr.field.IsLongOrDouble)
}
