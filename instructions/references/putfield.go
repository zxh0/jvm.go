package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Set field in object
type PutField struct {
	base.Index16Instruction
	field *heap.Field
}

func (instr *PutField) Execute(frame *rtda.Frame) {
	if instr.field == nil {
		cp := frame.GetConstantPool()
		kFieldRef := cp.GetConstantFieldRef(instr.Index)
		instr.field = kFieldRef.GetField(false)
	}

	val := frame.PopL(instr.field.IsLongOrDouble)
	ref := frame.PopRef()
	if ref == nil {
		frame.Thread.ThrowNPE()
		return
	}

	instr.field.PutValue(ref, val)
}
