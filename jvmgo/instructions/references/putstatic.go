package references

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

// Set static field in class
type PUT_STATIC struct {
	base.Index16Instruction
	field *heap.Field
}

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	if self.field == nil {
		cp := frame.Method().Class().ConstantPool()
		kFieldRef := cp.GetConstant(self.Index).(*heap.ConstantFieldref)
		self.field = kFieldRef.StaticField()
	}

	class := self.field.Class()
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread().InitClass(class)
		return
	}

	val := frame.OperandStack().PopField(self.field.IsLongOrDouble)
	self.field.PutStaticValue(val)
}
