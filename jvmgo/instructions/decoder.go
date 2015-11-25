package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
)

type Decoder base.BytecodeReader

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (self *Decoder) Decode(code []byte, pc int) (inst base.Instruction, nextPC int) {
	reader := (*base.BytecodeReader)(self)
	reader.Init(code, pc)

	opcode := reader.ReadUint8()
	inst = newInstruction(opcode)
	inst.FetchOperands(reader)
	nextPC = reader.PC()

	return
}
