package math

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Increment local variable by constant
type IINC struct {
	Index uint
	Const int32
}

func (instr *IINC) FetchOperands(reader *base.BytecodeReader) {
	instr.Index = uint(reader.ReadUint8())
	instr.Const = int32(reader.ReadInt8())
}

func (instr *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(instr.Index)
	val += instr.Const
	localVars.SetInt(instr.Index, val)
}
