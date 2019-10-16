package instructions

import (
	"github.com/zxh0/jvm.go/instructions/base"
)

func Decode(code []byte) []base.Instruction {
	reader := base.NewCodeReader(code)
	decoded := make([]base.Instruction, len(code))

	for reader.Position() < len(code) {
		decoded[reader.Position()] = decodeInstruction(reader)
	}

	return decoded
}

func decodeInstruction(reader *base.CodeReader) base.Instruction {
	opcode := reader.ReadUint8()
	instr := newInstruction(opcode)
	instr.FetchOperands(reader)
	return instr
}
