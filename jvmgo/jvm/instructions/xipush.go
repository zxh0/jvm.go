package instructions

import "github.com/zxh0/jvm.go/jvmgo/jvm/rtda"

// Push byte
type bipush struct {
	val int8
}

func (self *bipush) fetchOperands(decoder *InstructionDecoder) {
	self.val = decoder.readInt8()
}
func (self *bipush) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

// Push short
type sipush struct {
	val int16
}

func (self *sipush) fetchOperands(decoder *InstructionDecoder) {
	self.val = decoder.readInt16()
}
func (self *sipush) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
