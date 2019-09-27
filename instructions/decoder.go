package instructions

import (
	"github.com/zxh0/jvm.go/instructions/base"
)

func Decode(code []byte) []base.Instruction {
	reader := base.NewCodeReader(code)
	decoded := make([]base.Instruction, len(code))

	pc := 0
	for pc < len(code) {
		instr := decodeInstruction(reader)
		decoded[pc], pc = instr, reader.PC()
	}

	return decoded
}

func decodeInstruction(reader *base.CodeReader) base.Instruction {
	opcode := reader.ReadUint8()
	instr := newInstruction(opcode)
	instr.FetchOperands(reader)
	return instr
}
