package references

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

// Invoke instance method; dispatch based on class
type INVOKE_VIRTUAL struct {
	base.Index16Instruction
	kMethodRef   *heap.ConstantMethodref
	argSlotCount uint
}

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	if self.kMethodRef == nil {
		cp := frame.Method().ConstantPool()
		self.kMethodRef = cp.GetConstant(self.Index).(*heap.ConstantMethodref)
		self.argSlotCount = self.kMethodRef.ArgSlotCount()
	}

	stack := frame.OperandStack()
	ref := stack.TopRef(self.argSlotCount)
	if ref == nil {
		frame.Thread().ThrowNPE()
		return
	}

	method := self.kMethodRef.GetVirtualMethod(ref)
	frame.Thread().InvokeMethod(method)
}
