package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/rtda/linker"
)

// Fetch field from object
type GetField struct {
	base.Index16Instruction
	field *heap.Field
}

func (instr *GetField) Execute(frame *rtda.Frame) {
	if instr.field == nil {
		cp := frame.GetConstantPool()
		fieldRef := cp.GetConstantFieldRef(instr.Index)
		instr.field = linker.ResolveField(frame.GetBootLoader(), fieldRef, false)
	}

	ref := frame.PopRef()
	if ref == nil {
		frame.Thread.ThrowNPE()
		return
	}

	val := instr.field.GetValue(ref)
	frame.PushL(val, instr.field.IsLongOrDouble)
}
