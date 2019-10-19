package extended

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/instructions/control"
	"github.com/zxh0/jvm.go/instructions/loads"
	"github.com/zxh0/jvm.go/instructions/math"
	"github.com/zxh0/jvm.go/instructions/stores"
	"github.com/zxh0/jvm.go/rtda"
)

// Extend local variable index by additional bytes
type Wide struct {
	modifiedInstruction base.Instruction
}

func (instr *Wide) FetchOperands(reader *base.CodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15, 0x17, 0x19:
		inst := loads.NewLoad(false)
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0x16, 0x18:
		inst := loads.NewLoad(true)
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0x36, 0x38, 0x3a:
		inst := stores.NewStore(false)
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0x37, 0x39:
		inst := stores.NewStore(true)
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0xa9:
		inst := &control.RET{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0x84:
		inst := &math.IInc{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		instr.modifiedInstruction = inst
	}
}

func (instr *Wide) Execute(frame *rtda.Frame) {
	instr.modifiedInstruction.Execute(frame)
}
