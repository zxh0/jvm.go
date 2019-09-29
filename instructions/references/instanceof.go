package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Determine if object is of given type
type InstanceOf struct{ base.Index16Instruction }

func (instr *InstanceOf) Execute(frame *rtda.Frame) {
	ref := frame.PopRef()

	cp := frame.GetConstantPool()
	kClass := cp.GetConstant(instr.Index).(*heap.ConstantClass)
	class := kClass.Class()

	if ref == nil {
		frame.PushInt(0)
	} else if ref.IsInstanceOf(class) {
		frame.PushInt(1)
	} else {
		frame.PushInt(0)
	}
}
