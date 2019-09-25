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
type WIDE struct {
	modifiedInstruction base.Instruction
}

func (instr *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		inst := &loads.ILoad{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0x16:
		inst := &loads.LLoad{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0x17:
		inst := &loads.FLoad{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0x18:
		inst := &loads.DLoad{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0x19:
		inst := &loads.ALoad{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0x36:
		inst := &stores.IStore{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0x37:
		inst := &stores.LStore{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0x38:
		inst := &stores.FStore{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0x39:
		inst := &stores.DStore{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0x3a:
		inst := &stores.AStore{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0xa9:
		inst := &control.RET{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		instr.modifiedInstruction = inst
	}
}

func (instr *WIDE) Execute(frame *rtda.Frame) {
	instr.modifiedInstruction.Execute(frame)
}
