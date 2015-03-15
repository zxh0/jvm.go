package instructions

import (
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

// Determine if object is of given type
type instanceof struct{ Index16Instruction }

func (self *instanceof) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()

	cp := frame.ConstantPool()
	kClass := cp.GetConstant(self.index).(*rtc.ConstantClass)
	class := kClass.Class()

	if ref == nil {
		stack.PushInt(0)
	} else if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
