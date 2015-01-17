package instructions

type InstructionDecoder struct {
    pc          int
    bytecodes   []byte
}

func newInstructionDecoder(bytecodes []byte) (*InstructionDecoder) {
    return &InstructionDecoder{0, bytecodes}
}
