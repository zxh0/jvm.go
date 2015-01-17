package instructions

type Instruction interface {
    fetchOperands(bcr *BytecodeReader)
    execute()
}

func decode(bcr *BytecodeReader) (Instruction) {
    // todo
    return nil
}

func newInstruction(opcode byte) (Instruction) {
    switch opcode {
    case 0x32: return &aaload{}
    case 0x53: return &aastore{}
    default: panic("BAD opcode!")
    }
}
