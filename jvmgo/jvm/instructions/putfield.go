package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

// Set field in object
type putfield struct{ Index16Instruction }

func (self *putfield) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.Pop()
	ref := stack.PopRef()
	if ref == nil {
		frame.Thread().ThrowNPE()
		return
	}

	cp := frame.ConstantPool()
	kFieldRef := cp.GetConstant(self.index).(*rtc.ConstantFieldref)
	field := kFieldRef.InstanceField()

	field.PutValue(ref, val)
}

// Set static field in class
type putstatic struct{ Index16Instruction }

func (self *putstatic) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	kFieldRef := cp.GetConstant(self.index).(*rtc.ConstantFieldref)
	field := kFieldRef.StaticField()

	class := field.Class()
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread().InitClass(class)
		return
	}

	val := frame.OperandStack().Pop()
	field.PutStaticValue(val)
}
