package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Invoke interface method
type INVOKE_INTERFACE struct {
	index uint
	// count uint8
	// zero uint8

	// optimization
	kMethodRef   *heap.ConstantInterfaceMethodref
	argSlotCount uint
}

func (instr *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	instr.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count
	reader.ReadUint8() // must be 0
}

func (instr *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	if instr.kMethodRef == nil {
		cp := frame.Method().ConstantPool()
		instr.kMethodRef = cp.GetConstant(instr.index).(*heap.ConstantInterfaceMethodref)
		instr.argSlotCount = instr.kMethodRef.ArgSlotCount()
	}

	stack := frame.OperandStack()
	ref := stack.TopRef(instr.argSlotCount)
	if ref == nil {
		panic("NPE") // todo
	}

	method := instr.kMethodRef.FindInterfaceMethod(ref)
	frame.Thread().InvokeMethod(method)
}
