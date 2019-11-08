package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/rtda/linker"
	"github.com/zxh0/jvm.go/vm"
)

// Invoke interface method
type InvokeInterface struct {
	index uint
	// count uint8
	// zero uint8

	// optimization
	methodRef *heap.ConstantMethodRef
}

func (instr *InvokeInterface) FetchOperands(reader *base.CodeReader) {
	instr.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count
	reader.ReadUint8() // must be 0
}

func (instr *InvokeInterface) Execute(frame *rtda.Frame) {
	if instr.methodRef == nil {
		cp := frame.GetConstantPool()
		methodRef := cp.GetConstant(instr.index).(*heap.ConstantMethodRef)
		if !methodRef.IsInterface {
			panic(vm.NewIncompatibleClassChangeError(methodRef.String()))
		}

		method := linker.ResolveMethod(frame.GetBootLoader(), methodRef)
		if method.IsStatic() || method.IsConstructor() {
			panic(vm.NewIncompatibleClassChangeError(method.String()))
		}

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
