package instructions

type Instruction interface {

}

func decode(bytecodes []byte, pc uint) (Instruction) {
    opcode := bytecodes[pc]
    

    // todo
    return nil
}

func newInstruction(opcode byte) (Instruction) {
    switch opcode {
    case 0x32: return &aaload{}
    default: panic("BAD opcode!")
    }
}
