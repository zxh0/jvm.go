package constants

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Push byte
type bipush struct {
	val int8
}

func (self *bipush) fetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *bipush) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

// Push short
type sipush struct {
	val int16
}

func (self *sipush) fetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *sipush) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
