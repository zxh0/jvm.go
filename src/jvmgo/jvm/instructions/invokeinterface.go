package instructions

import (
	//"fmt"
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

// Invoke interface method
type invokeinterface struct {
	index uint16
	count uint8 // unused
	// 0
}

func (self *invokeinterface) fetchOperands(decoder *InstructionDecoder) {
	self.index = decoder.readUint16()
	self.count = decoder.readUint8()
	decoder.readUint8() // must be 0
}

func (self *invokeinterface) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(uint(self.index))
	cMethodRef := c.(*rtc.ConstantInterfaceMethodref)
	ref := stack.Top(cMethodRef.ArgCount()).(*rtc.Obj)

	if ref == nil {
		panic("NPE") // todo
	}

	method := cMethodRef.VirtualMethod(ref)
	thread.InvokeMethod(method)
}
