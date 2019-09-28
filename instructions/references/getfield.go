package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Fetch field from object
type GetField struct {
	base.Index16Instruction
	field *heap.Field
}

func (instr *GetField) Execute(frame *rtda.Frame) {
	if instr.field == nil {
		cp := frame.Method().Class().ConstantPool()
		kFieldRef := cp.GetConstant(instr.Index).(*heap.ConstantFieldRef)
		instr.field = kFieldRef.InstanceField()
	}

	ref := frame.PopRef()
	if ref == nil {
		frame.Thread().ThrowNPE()
		return
	}

	val := instr.field.GetValue(ref)
	frame.PushL(val, instr.field.IsLongOrDouble)
}
