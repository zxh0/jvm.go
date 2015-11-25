package references

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

// Check whether object is of given type
type checkcast struct {
	base.Index16Instruction
	class *rtc.Class
}

func (self *checkcast) Execute(frame *rtda.Frame) {
	if self.class == nil {
		cp := frame.Method().Class().ConstantPool()
		kClass := cp.GetConstant(self.Index).(*rtc.ConstantClass)
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
