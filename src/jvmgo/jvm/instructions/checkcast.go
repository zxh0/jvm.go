package instructions

import (
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

// Check whether object is of given type
type checkcast struct{ Index16Instruction }

func (self *checkcast) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)

	if ref == nil {
		return
	}

	cp := frame.ConstantPool()
	kClass := cp.GetConstant(self.index).(*rtc.ConstantClass)
	class := kClass.Class()
	if class.InitializationNotStarted() {
		// todo: init class
		panic("class not initialized!" + class.Name())
	}

	if !ref.IsInstanceOf(class) {
		frame.Thread().ThrowClassCastException(ref.Class(), class)
	}
}
