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
		cp := frame.Method().Class().ConstantPool()
		kFieldRef := cp.GetConstant(instr.Index).(*heap.ConstantFieldref)
		instr.field = kFieldRef.InstanceField()
	}

	stack := frame.OperandStack()
	val := stack.PopField(instr.field.IsLongOrDouble)
	ref := stack.PopRef()
	if ref == nil {
		frame.Thread().ThrowNPE()
		return
	}

	instr.field.PutValue(ref, val)
}
