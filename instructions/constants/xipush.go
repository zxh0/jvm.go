package constants

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Push byte
type BIPush struct {
	val int8
}

func (instr *BIPush) FetchOperands(reader *base.BytecodeReader) {
	instr.val = reader.ReadInt8()
}
func (instr *BIPush) Execute(frame *rtda.Frame) {
	i := int32(instr.val)
	frame.OperandStack().PushInt(i)
}

// Push short
type SIPush struct {
	val int16
}

func (instr *SIPush) FetchOperands(reader *base.BytecodeReader) {
	instr.val = reader.ReadInt16()
}
func (instr *SIPush) Execute(frame *rtda.Frame) {
	i := int32(instr.val)
	frame.OperandStack().PushInt(i)
}
