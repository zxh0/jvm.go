package references

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
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

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count
	reader.ReadUint8() // must be 0
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	if self.kMethodRef == nil {
		cp := frame.Method().ConstantPool()
		self.kMethodRef = cp.GetConstant(self.index).(*heap.ConstantInterfaceMethodref)
		self.argSlotCount = self.kMethodRef.ArgSlotCount()
	}

	stack := frame.OperandStack()
	ref := stack.TopRef(self.argSlotCount)
	if ref == nil {
		panic("NPE") // todo
	}

	method := self.kMethodRef.FindInterfaceMethod(ref)
	frame.Thread().InvokeMethod(method)
}
