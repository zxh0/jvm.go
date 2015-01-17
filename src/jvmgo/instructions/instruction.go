package instructions

type Instruction interface {
    fetchOperands(bcr *BytecodeReader)
    execute()
}

func decode(bcr *BytecodeReader) (Instruction) {
    opcode := bcr.readOpcode()
    instruction := newInstruction(opcode)
    instruction.fetchOperands(bcr)
    return instruction
}

func newInstruction(opcode byte) (Instruction) {
    switch opcode {
    case 0x32: return &aaload{}
    case 0x53: return &aastore{}
    default: panic("BAD opcode!")
    }
}
