package instructions

import (
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

// Create new object
type new_ struct{ Index16Instruction }

func (self *new_) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	kClass := cp.GetConstant(self.index).(*rtc.ConstantClass)
	class := kClass.Class()

	if class.InitializationNotStarted() {
		frame.RevertNextPC() // undo new
		rtda.InitClass(class, frame.Thread())
	} else {
		ref := class.NewObj()
		frame.OperandStack().PushRef(ref)
	}
}
