package extended

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Extend local variable index by additional bytes
type wide struct {
	modifiedInstruction base.Instruction
}

func (self *wide) fetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		inst := &iload{}
		inst.index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x16:
		inst := &lload{}
		inst.index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x17:
		inst := &fload{}
		inst.index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x18:
		inst := &dload{}
		inst.index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x19:
		inst := &aload{}
		inst.index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x36:
		inst := &istore{}
		inst.index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x37:
		inst := &lstore{}
		inst.index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x38:
		inst := &fstore{}
		inst.index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x39:
		inst := &dstore{}
		inst.index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x3a:
		inst := &astore{}
		inst.index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0xa9:
		inst := &ret{}
		inst.index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x84:
		inst := &iinc{}
		inst.index = uint(reader.ReadUint16())
		inst._const = int32(reader.ReadInt16())
		self.modifiedInstruction = inst
	}
}

func (self *wide) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}
