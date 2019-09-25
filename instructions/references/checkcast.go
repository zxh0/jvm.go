package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Check whether object is of given type
type CHECK_CAST struct {
	base.Index16Instruction
	class *heap.Class
}

func (instr *CHECK_CAST) Execute(frame *rtda.Frame) {
	if instr.class == nil {
		cp := frame.Method().Class().ConstantPool()
		kClass := cp.GetConstant(instr.Index).(*heap.ConstantClass)
		instr.class = kClass.Class()
	}

	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)

	if ref == nil {
		return
	}

	if !ref.IsInstanceOf(instr.class) {
		frame.Thread().ThrowClassCastException(ref.Class(), instr.class)
	}
}
