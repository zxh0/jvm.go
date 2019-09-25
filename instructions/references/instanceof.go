package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Determine if object is of given type
type INSTANCE_OF struct{ base.Index16Instruction }

func (instr *INSTANCE_OF) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()

	cp := frame.ConstantPool()
	kClass := cp.GetConstant(instr.Index).(*heap.ConstantClass)
	class := kClass.Class()

	if ref == nil {
		stack.PushInt(0)
	} else if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
