package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Invoke instance method; dispatch based on class
type InvokeVirtual struct {
	base.Index16Instruction
	kMethodRef   *heap.ConstantMethodRef
	argSlotCount uint
}

func (instr *InvokeVirtual) Execute(frame *rtda.Frame) {
	if instr.kMethodRef == nil {
		cp := frame.GetConstantPool()
		instr.kMethodRef = cp.GetConstant(instr.Index).(*heap.ConstantMethodRef)
		instr.argSlotCount = instr.kMethodRef.ParamSlotCount
	}

	ref := frame.TopRef(instr.argSlotCount)
	if ref == nil {
		frame.Thread.ThrowNPE()
		return
	}

	method := instr.kMethodRef.GetVirtualMethod(ref)
	frame.Thread.InvokeMethod(method)
}
