package instructions

// todo
func Decode(bcr *BytecodeReader) (uint8, Instruction) {
    opcode := bcr.readUint8()
    instruction := newInstruction(opcode)
    instruction.fetchOperands(bcr)
    return opcode, instruction
}
