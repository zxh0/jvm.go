package instructions

import "jvmgo/rtda"

type Instruction interface {
    fetchOperands(bcr *BytecodeReader)
    execute(thread *rtda.Thread)
}

func decode(bcr *BytecodeReader) (Instruction) {
    opcode := bcr.readUint8()
    instruction := newInstruction(opcode)
    instruction.fetchOperands(bcr)
    return instruction
}

func newInstruction(opcode byte) (Instruction) {
    switch opcode {
    case 0x01: return &aconst_null{}
    case 0x19: return &aload{}
    case 0x2a: return &aload_0{}
    case 0x2b: return &aload_1{}
    case 0x2c: return &aload_2{}
    case 0x2d: return &aload_3{}
    case 0x32: return &aaload{}
    case 0x3a: return &astore{}
    case 0x4b: return &astore_0{}
    case 0x4c: return &astore_1{}
    case 0x4d: return &astore_2{}
    case 0x4e: return &astore_3{}
    case 0x53: return &aastore{}
    case 0xb0: return &areturn{}
    case 0xbd: return &anewarray{}
    case 0xbe: return &arraylength{}
    default: panic("BAD opcode!")
    }
}
