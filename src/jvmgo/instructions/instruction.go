package instructions

type Instruction interface {
    fetchOperands(bcr *BytecodeReader)
}

func decode(bcr *BytecodeReader) (Instruction) {
    // todo
    return nil
}

func newInstruction(opcode byte) (Instruction) {
    switch opcode {
    case 0x32: return &aaload{}
    default: panic("BAD opcode!")
    }
}
