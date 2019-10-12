package instructions

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/instructions/control"
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
	pcMap := make(map[int]int) // pc -> compactPC

	compactPC := 0
	for pc, instr := range decoded {
		if instr != nil {
			pcMap[pc] = compactPC
			compactPC++
		}
	}

	compactPC = 0
	for pc, instr := range decoded {
		if instr != nil {
			if bi, ok := instr.(_BI); ok {
				bi.SetOffset(adjustOffset(pc, bi.GetOffset(), pcMap))
			} else if tableSwitch, ok := instr.(*control.TableSwitch); ok {
				fixTableSwitchOffsets(tableSwitch, pc, pcMap)
			} else if lookupSwitch, ok := instr.(*control.LookupSwitch); ok {
				fixLookupSwitchOffsets(lookupSwitch, pc, pcMap)
			}

			decoded[compactPC] = instr
			compactPC++
		}
	}

	return decoded[:compactPC]
}

func fixTableSwitchOffsets(instr *control.TableSwitch, pc int, pcMap map[int]int) {
	instr.DefaultOffset = int32(adjustOffset(pc, int(instr.DefaultOffset), pcMap))
	for i, offset := range instr.JumpOffsets {
		instr.JumpOffsets[i] = int32(adjustOffset(pc, int(offset), pcMap))
	}
}

func fixLookupSwitchOffsets(instr *control.LookupSwitch, pc int, pcMap map[int]int) {
	instr.DefaultOffset = int32(adjustOffset(pc, int(instr.DefaultOffset), pcMap))
	for i, offset := range instr.MatchOffsets {
		if i%2 == 1 {
			instr.MatchOffsets[i] = int32(adjustOffset(pc, int(offset), pcMap))
		}
	}
}

func adjustOffset(pc, offset int, pcMap map[int]int) int {
	return pcMap[pc+offset] - pcMap[pc]
}
