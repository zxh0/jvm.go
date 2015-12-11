package references

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

// Check whether object is of given type
type CHECK_CAST struct {
	base.Index16Instruction
	class *heap.Class
}

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	if self.class == nil {
		cp := frame.Method().Class().ConstantPool()
		kClass := cp.GetConstant(self.Index).(*heap.ConstantClass)
		self.class = kClass.Class()
	}

	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)

	if ref == nil {
		return
	}

	if !ref.IsInstanceOf(self.class) {
		frame.Thread().ThrowClassCastException(ref.Class(), self.class)
	}
}
