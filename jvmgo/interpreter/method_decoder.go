package interpreter

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions"
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
)

var decoder = instructions.NewDecoder()

func decodeMethod(code []byte) []base.Instruction {
	insts := make([]base.Instruction, len(code))

	pc := 0
	for pc < len(code) {
		inst, nextPC := decoder.Decode(code, pc)
		insts[pc], pc = inst, nextPC
	}

	return insts
}
