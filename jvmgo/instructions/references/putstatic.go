package references

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

// Set static field in class
type putstatic struct {
	base.Index16Instruction
	field *rtc.Field
}

func (self *putstatic) Execute(frame *rtda.Frame) {
	if self.field == nil {
		cp := frame.Method().Class().ConstantPool()
		kFieldRef := cp.GetConstant(self.Index).(*rtc.ConstantFieldref)
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
