package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/rtda/linker"
)

// Invoke instance method; dispatch based on class
type InvokeVirtual struct {
	base.Index16Instruction
	methodRef *heap.ConstantMethodRef
}

func (instr *InvokeVirtual) Execute(frame *rtda.Frame) {
	if instr.methodRef == nil {
		cp := frame.GetConstantPool()
		methodRef := cp.GetConstant(instr.Index).(*heap.ConstantMethodRef)
		linker.ResolveMethod(frame.GetBootLoader(), methodRef)

		instr.methodRef = methodRef
	}

	obj := frame.TopRef(instr.methodRef.ResolvedMethod.ParamSlotCount - 1)
	if obj == nil {
		frame.Thread.ThrowNPE()
		return
	}

	method := linker.SelectMethod(obj, instr.methodRef)
	frame.Thread.InvokeMethod(method)
}
