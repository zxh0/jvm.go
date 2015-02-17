package instructions

import (
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

// Create new array
type newarray struct {
	atype uint8
}

func (self *newarray) fetchOperands(decoder *InstructionDecoder) {
	self.atype = decoder.readUint8()
}
func (self *newarray) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	count := uint(stack.PopInt())
	classLoader := frame.Method().Class().ClassLoader()
	arr := rtc.NewPrimitiveArray(self.atype, count, classLoader)
	stack.PushRef(arr)
}

// Create new array of reference
type anewarray struct{ Index16Instruction }

func (self *anewarray) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	cClass := cp.GetConstant(self.index).(*rtc.ConstantClass)
	componentClass := cClass.Class()

	if componentClass.InitializationNotStarted() {
		thread := frame.Thread()
		frame.SetNextPC(thread.PC()) // undo anewarray
		thread.InitClass(componentClass)
		return
	}

	stack := frame.OperandStack()
	count := uint(stack.PopInt())
	arr := rtc.NewRefArray(componentClass, count)
	stack.PushRef(arr)
}

// Create new multidimensional array
type multianewarray struct {
	index      uint16
	dimensions uint8
}

func (self *multianewarray) fetchOperands(decoder *InstructionDecoder) {
	self.index = decoder.readUint16()
	self.dimensions = decoder.readUint8()
}
func (self *multianewarray) Execute(frame *rtda.Frame) {
	// todo
	panic("todo multianewarray")
}
