package constants

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Push byte
type BIPUSH struct {
	val int8
}

func (instr *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	instr.val = reader.ReadInt8()
}
func (instr *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(instr.val)
	frame.OperandStack().PushInt(i)
}

// Push short
type SIPUSH struct {
	val int16
}

func (instr *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	instr.val = reader.ReadInt16()
}
func (instr *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(instr.val)
	frame.OperandStack().PushInt(i)
}
