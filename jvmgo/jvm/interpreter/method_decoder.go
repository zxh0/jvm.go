package interpreter

import (
	"github.com/zxh0/jvm.go/jvmgo/jvm/instructions"
)

var decoder = instructions.NewDecoder()

func decodeMethod(code []byte) []instructions.Instruction {
	insts := make([]instructions.Instruction, len(code))

	pc := 0
	for pc < len(code) {
		inst, nextPC := decoder.Decode(code, pc)
		insts[pc], pc = inst, nextPC
	}

	return insts
}
