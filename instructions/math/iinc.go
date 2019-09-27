package math

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Increment local variable by constant
type IInc struct {
	Index uint
	Const int32
}

func (instr *IInc) FetchOperands(reader *base.CodeReader) {
	instr.Index = uint(reader.ReadUint8())
	instr.Const = int32(reader.ReadInt8())
}

func (instr *IInc) Execute(frame *rtda.Frame) {
	val := frame.GetIntVar(instr.Index)
	val += instr.Const
	frame.SetIntVar(instr.Index, val)
}
