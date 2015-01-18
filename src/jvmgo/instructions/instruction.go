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
    case 0x0e: return &dconst_0{}
    case 0x0f: return &dconst_1{}
    case 0x10: return &bipush{}
    case 0x18: return &dload{}
    case 0x19: return &aload{}
    case 0x26: return &dload_0{}
    case 0x27: return &dload_1{}
    case 0x28: return &dload_2{}
    case 0x29: return &dload_3{}
    case 0x2a: return &aload_0{}
    case 0x2b: return &aload_1{}
    case 0x2c: return &aload_2{}
    case 0x2d: return &aload_3{}
    case 0x31: return &daload{}
    case 0x32: return &aaload{}
    case 0x33: return &baload{}
    case 0x34: return &caload{}
    case 0x38: return &dstore{}
    case 0x3a: return &astore{}
    case 0x47: return &dstore_0{}
    case 0x48: return &dstore_1{}
    case 0x49: return &dstore_2{}
    case 0x4a: return &dstore_3{}
    case 0x4b: return &astore_0{}
    case 0x4c: return &astore_1{}
    case 0x4d: return &astore_2{}
    case 0x4e: return &astore_3{}
    case 0x52: return &dastore{}
    case 0x53: return &aastore{}
    case 0x54: return &bastore{}
    case 0x55: return &castore{}
    case 0x59: return &dup{}
    case 0x5a: return &dup_x1{}
    case 0x5b: return &dup_x2{}
    case 0x60: return &iadd{}
    case 0x61: return &ladd{}
    case 0x62: return &fadd{}
    case 0x63: return &dadd{}
    case 0x64: return &isub{}
    case 0x65: return &lsub{}
    case 0x66: return &fsub{}
    case 0x67: return &dsub{}
    case 0x68: return &imul{}
    case 0x69: return &lmul{}
    case 0x6a: return &fmul{}
    case 0x6b: return &dmul{}
    case 0x6c: return &idiv{}
    case 0x6d: return &ldiv{}
    case 0x6e: return &fdiv{}
    case 0x6f: return &ddiv{}
    case 0x70: return &irem{}
    case 0x71: return &lrem{}
    case 0x72: return &frem{}
    case 0x73: return &drem{}
    case 0x74: return &ineg{}
    case 0x75: return &lneg{}
    case 0x76: return &fneg{}
    case 0x77: return &dneg{}
    case 0x8e: return &d2i{}
    case 0x8f: return &d2l{}
    case 0x90: return &d2f{}
    case 0x97: return &dcmpl{}
    case 0x98: return &dcmpg{}
    case 0xaf: return &dreturn{}
    case 0xb0: return &areturn{}
    case 0xbd: return &anewarray{}
    case 0xbe: return &arraylength{}
    case 0xbf: return &athrow{}
    case 0xc0: return &checkcast{}
    default: panic("BAD opcode!")
    }
}
