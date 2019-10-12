package instructions

import (
	"github.com/zxh0/jvm.go/instructions/base"
)

// BranchInstruction
type _BI interface {
	GetOffset() int
	SetOffset(offset int)
}

func Decode(code []byte, compact bool) []base.Instruction {
	reader := base.NewCodeReader(code)
	decoded := make([]base.Instruction, len(code))

	for reader.Position() < len(code) {
		decoded[reader.Position()] = decodeInstruction(reader)
	}

	if compact {
		return compactInstructions(decoded)
	} else {
		return decoded
	}
}

func decodeInstruction(reader *base.CodeReader) base.Instruction {
	opcode := reader.ReadUint8()
	instr := newInstruction(opcode)
	instr.FetchOperands(reader)
	return instr
}

func compactInstructions(decoded []base.Instruction) []base.Instruction {
	pcMap := make(map[int]int) // bytecodePC -> instrPC

	instrPC := 0
	for bytecodePC, instr := range decoded {
		if instr != nil {
			pcMap[bytecodePC] = instrPC
			instrPC++
		}
	}

	instrPC = 0
	for bytecodePC, instr := range decoded {
		if instr != nil {
			if bi, ok := instr.(_BI); ok {
				bytecodeTarget := bytecodePC + bi.GetOffset()
				if instrTarget, found := pcMap[bytecodeTarget]; found {
					bi.SetOffset(instrTarget - instrPC)
				} else {
					panic("!!!")
				}
			}
			decoded[instrPC] = instr
			instrPC++
		}
	}

	return decoded[:instrPC]
}
