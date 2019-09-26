package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Check whether object is of given type
type CheckCast struct {
	base.Index16Instruction
	class *heap.Class
}

func (instr *CheckCast) Execute(frame *rtda.Frame) {
	if instr.class == nil {
		cp := frame.Method().Class().ConstantPool()
		kClass := cp.GetConstant(instr.Index).(*heap.ConstantClass)
		instr.class = kClass.Class()
	}

	ref := frame.PopRef()
	frame.PushRef(ref)

	if ref == nil {
		return
	}

	if !ref.IsInstanceOf(instr.class) {
		frame.Thread().ThrowClassCastException(ref.Class(), instr.class)
	}
}
