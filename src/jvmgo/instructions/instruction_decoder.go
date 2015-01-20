package instructions

var (
    _aload_0 = &aload_0{}
    _aload_1 = &aload_1{}
    _aload_2 = &aload_2{}
    _aload_3 = &aload_3{}
    _arraylength = &arraylength{}
    _astore_0 = &astore_0{}
    _astore_1 = &astore_1{}
    _astore_2 = &astore_2{}
    _astore_3 = &astore_3{}
    _athrow = &athrow{}
    _d2f = &d2f{}
    _d2i = &d2i{}
    _d2l = &d2l{}
    _dcmpg = &dcmpg{}
    _dcmpl = &dcmpl{}
    _dload_0 = &dload_0{}
    _dload_1 = &dload_1{}
    _dload_2 = &dload_2{}
    _dload_3 = &dload_3{}
    _dstore_0 = &dstore_0{}
    _dstore_1 = &dstore_1{}
    _dstore_2 = &dstore_2{}
    _dstore_3 = &dstore_3{}
    _dup = &dup{}
    _dup_x1 = &dup_x1{}
    _dup_x2 = &dup_x2{}
    _dup2 = &dup2{}
    _dup2_x1 = &dup2_x1{}
    _dup2_x2 = &dup2_x2{}
    _f2d = &f2d{}
    _f2i = &f2i{}
    _f2l = &f2l{}
    _fcmpg = &fcmpg{}
    _fcmpl = &fcmpl{}
    _fload_0 = &fload_0{}
    _fload_1 = &fload_1{}
    _fload_2 = &fload_2{}
    _fload_3 = &fload_3{}
    _fstore_0 = &fstore_0{}
    _fstore_1 = &fstore_1{}
    _fstore_2 = &fstore_2{}
    _fstore_3 = &fstore_3{}
)

func Decode(bcr *BytecodeReader) (Instruction) {
    opcode := bcr.readUint8()
    instruction := newInstruction(opcode)
    instruction.fetchOperands(bcr)
    return instruction
}

func newInstruction(opcode byte) (Instruction) {
    switch opcode {
    case 0x00: return &nop{}
    case 0x01: return &aconst_null{}
    case 0x02: return &iconst_m1{}
    case 0x03: return &iconst_0{}
    case 0x04: return &iconst_1{}
    case 0x05: return &iconst_2{}
    case 0x06: return &iconst_3{}
    case 0x07: return &iconst_4{}
    case 0x08: return &iconst_5{}
    case 0x09: return &lconst_0{}
    case 0x0a: return &lconst_1{}
    case 0x0b: return &fconst_0{}
    case 0x0c: return &fconst_1{}
    case 0x0d: return &fconst_2{}
    case 0x0e: return &dconst_0{}
    case 0x0f: return &dconst_1{}
    case 0x10: return &bipush{}
    case 0x11: return &sipush{}
    case 0x12: return &ldc{}
    case 0x13: return &ldc_w{}
    case 0x14: return &ldc2_w{}
    case 0x15: return &iload{}
    case 0x16: return &lload{}
    case 0x17: return &fload{}
    case 0x18: return &dload{}
    case 0x19: return &aload{}
    case 0x1a: return &iload_0{}
    case 0x1b: return &iload_1{}
    case 0x1c: return &iload_2{}
    case 0x1d: return &iload_3{}
    case 0x1e: return &lload_0{}
    case 0x1f: return &lload_1{}
    case 0x20: return &lload_2{}
    case 0x21: return &lload_3{}
    case 0x22: return _fload_0
    case 0x23: return _fload_1
    case 0x24: return _fload_2
    case 0x25: return _fload_3
    case 0x26: return _dload_0
    case 0x27: return _dload_1
    case 0x28: return _dload_2
    case 0x29: return _dload_3
    case 0x2a: return _aload_0
    case 0x2b: return _aload_1
    case 0x2c: return _aload_2
    case 0x2d: return _aload_3
    case 0x2e: return &iaload{}
    case 0x2f: return &laload{}
    case 0x30: return &faload{}
    case 0x31: return &daload{}
    case 0x32: return &aaload{}
    case 0x33: return &baload{}
    case 0x34: return &caload{}
    case 0x35: return &saload{}
    case 0x36: return &istore{}
    case 0x37: return &lstore{}
    case 0x38: return &fstore{}
    case 0x39: return &dstore{}
    case 0x3a: return &astore{}
    case 0x3b: return &istore_0{}
    case 0x3c: return &istore_1{}
    case 0x3d: return &istore_2{}
    case 0x3e: return &istore_3{}
    case 0x3f: return &lstore_0{}
    case 0x40: return &lstore_1{}
    case 0x41: return &lstore_2{}
    case 0x42: return &lstore_3{}
    case 0x43: return _fstore_0
    case 0x44: return _fstore_1
    case 0x45: return _fstore_2
    case 0x46: return _fstore_3
    case 0x47: return _dstore_0
    case 0x48: return _dstore_1
    case 0x49: return _dstore_2
    case 0x4a: return _dstore_3
    case 0x4b: return _astore_0
    case 0x4c: return _astore_1
    case 0x4d: return _astore_2
    case 0x4e: return _astore_3
    case 0x4f: return &iastore{}
    case 0x50: return &lastore{}
    case 0x51: return &fastore{}
    case 0x52: return &dastore{}
    case 0x53: return &aastore{}
    case 0x54: return &bastore{}
    case 0x55: return &castore{}
    case 0x56: return &sastore{}
    case 0x57: return &pop{}
    case 0x58: return &pop2{}
    case 0x59: return _dup
    case 0x5a: return _dup_x1
    case 0x5b: return _dup_x2
    case 0x5c: return _dup2
    case 0x5d: return _dup2_x1
    case 0x5e: return _dup2_x2
    case 0x5f: return &swap{}
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
    case 0x78: return &ishl{}
    case 0x79: return &lshl{}
    case 0x7a: return &ishr{}
    case 0x7b: return &lshr{}
    case 0x7c: return &iushr{}
    case 0x7d: return &lushr{}
    case 0x7e: return &iand{}
    case 0x7f: return &land{}
    case 0x80: return &ior{}
    case 0x81: return &lor{}
    case 0x82: return &ixor{}
    case 0x83: return &lxor{}
    case 0x84: return &iinc{}
    case 0x85: return &i2l{}
    case 0x86: return &i2f{}
    case 0x87: return &i2d{}
    case 0x88: return &l2i{}
    case 0x89: return &l2f{}
    case 0x8a: return &l2d{}
    case 0x8b: return _f2i
    case 0x8c: return _f2l
    case 0x8d: return _f2d
    case 0x8e: return _d2i
    case 0x8f: return _d2l
    case 0x90: return _d2f
    case 0x91: return &i2b{}
    case 0x92: return &i2c{}
    case 0x93: return &i2s{}
    case 0x94: return &lcmp{}
    case 0x95: return _fcmpl
    case 0x96: return _fcmpg
    case 0x97: return _dcmpl
    case 0x98: return _dcmpg
    case 0x99: return &ifeq{}
    case 0x9a: return &ifne{}
    case 0x9b: return &iflt{}
    case 0x9c: return &ifge{}
    case 0x9d: return &ifgt{}
    case 0x9e: return &ifle{}
    case 0x9f: return &if_icmpeq{}
    case 0xa0: return &if_icmpne{}
    case 0xa1: return &if_icmplt{}
    case 0xa2: return &if_icmpge{}
    case 0xa3: return &if_icmpgt{}
    case 0xa4: return &if_icmple{}
    case 0xa5: return &if_acmpeq{}
    case 0xa6: return &if_acmpne{}
    case 0xa7: return &_goto{}
  //case 0xa8: return &jsr{}
  //case 0xa9: return &ret{}
    case 0xaa: return &tableswitch{}
    case 0xab: return &lookupswitch{}
    case 0xac: return &ireturn{}
    case 0xad: return &lreturn{}
    case 0xae: return &freturn{}
    case 0xaf: return &dreturn{}
    case 0xb0: return &areturn{}
    case 0xb1: return &_return{}
    case 0xb2: return &getstatic{}
    case 0xb3: return &putstatic{}
    case 0xb4: return &getfield{}
    case 0xb5: return &putfield{}
    case 0xb6: return &invokevirtual{}
    case 0xb7: return &invokespecial{}
    case 0xb8: return &invokestatic{}
    case 0xb9: return &invokeinterface{}
    case 0xba: return &invokedynamic{}
    case 0xbb: return &_new{}
    case 0xbc: return &newarray{}
    case 0xbd: return &anewarray{}
    case 0xbe: return _arraylength
    case 0xbf: return _athrow
    case 0xc0: return &checkcast{}
    case 0xc1: return &instanceof{}
    case 0xc2: return &monitorenter{}
    case 0xc3: return &monitorexit{}
    case 0xc5: return &multianewarray{}
    case 0xc6: return &ifnull{}
    case 0xc7: return &ifnonnull{}
    case 0xc8: return &goto_w{}
  //case 0xc9: return &jsr_w{}
  //case 0xca: return &breakpoint{}
  //case 0xfe: return &impdep1{}
  //case 0xff: return &impdep2{}
    // todo
    default: panic("BAD opcode!")
    }
}
