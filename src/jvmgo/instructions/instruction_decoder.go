package instructions

// NoOperandsInstructions
var (
    _nop = &nop{}
    _aconst_null = &aconst_null{}
    _iconst_m1 = &iconst_m1{}
    _iconst_0 = &iconst_0{}
    _iconst_1 = &iconst_1{}
    _iconst_2 = &iconst_2{}
    _iconst_3 = &iconst_3{}
    _iconst_4 = &iconst_4{}
    _iconst_5 = &iconst_5{}
    _lconst_0 = &lconst_0{}
    _lconst_1 = &lconst_1{}
    _fconst_0 = &fconst_0{}
    _fconst_1 = &fconst_1{}
    _fconst_2 = &fconst_2{}
    _dconst_0 = &dconst_0{}
    _dconst_1 = &dconst_1{}
    _iload_0 = &iload_0{}
    _iload_1 = &iload_1{}
    _iload_2 = &iload_2{}
    _iload_3 = &iload_3{}
    _lload_0 = &lload_0{}
    _lload_1 = &lload_1{}
    _lload_2 = &lload_2{}
    _lload_3 = &lload_3{}
    _fload_0 = &fload_0{}
    _fload_1 = &fload_1{}
    _fload_2 = &fload_2{}
    _fload_3 = &fload_3{}
    _dload_0 = &dload_0{}
    _dload_1 = &dload_1{}
    _dload_2 = &dload_2{}
    _dload_3 = &dload_3{}
    _aload_0 = &aload_0{}
    _aload_1 = &aload_1{}
    _aload_2 = &aload_2{}
    _aload_3 = &aload_3{}
    _iaload = &iaload{}
    _laload = &laload{}
    _faload = &faload{}
    _daload = &daload{}
    _aaload = &aaload{}
    _baload = &baload{}
    _caload = &caload{}
    _saload = &saload{}
    _istore_0 = &istore_0{}
    _istore_1 = &istore_1{}
    _istore_2 = &istore_2{}
    _istore_3 = &istore_3{}
    _lstore_0 = &lstore_0{}
    _lstore_1 = &lstore_1{}
    _lstore_2 = &lstore_2{}
    _lstore_3 = &lstore_3{}
    _fstore_0 = &fstore_0{}
    _fstore_1 = &fstore_1{}
    _fstore_2 = &fstore_2{}
    _fstore_3 = &fstore_3{}
    _dstore_0 = &dstore_0{}
    _dstore_1 = &dstore_1{}
    _dstore_2 = &dstore_2{}
    _dstore_3 = &dstore_3{}
    _astore_0 = &astore_0{}
    _astore_1 = &astore_1{}
    _astore_2 = &astore_2{}
    _astore_3 = &astore_3{}
    _iastore = &iastore{}
    _lastore = &lastore{}
    _fastore = &fastore{}
    _dastore = &dastore{}
    _aastore = &aastore{}
    _bastore = &bastore{}
    _castore = &castore{}
    _sastore = &sastore{}
    _pop = &pop{}
    _pop2 = &pop2{}
    _dup = &dup{}
    _dup_x1 = &dup_x1{}
    _dup_x2 = &dup_x2{}
    _dup2 = &dup2{}
    _dup2_x1 = &dup2_x1{}
    _dup2_x2 = &dup2_x2{}
    _swap = &swap{}
    _iadd = &iadd{}
    _ladd = &ladd{}
    _fadd = &fadd{}
    _dadd = &dadd{}
    _idiv = &idiv{}
    _ldiv = &ldiv{}
    _fdiv = &fdiv{}
    _ddiv = &ddiv{}
    _imul = &imul{}
    _lmul = &lmul{}
    _fmul = &fmul{}
    _dmul = &dmul{}
    _ineg = &ineg{}
    _lneg = &lneg{}
    _fneg = &fneg{}
    _dneg = &dneg{}
    _irem = &irem{}
    _lrem = &lrem{}
    _frem = &frem{}
    _drem = &drem{}
    _isub = &isub{}
    _lsub = &lsub{}
    _fsub = &fsub{}
    _dsub = &dsub{}
    _iand = &iand{}
    _land = &land{}
    _ior = &ior{}
    _lor = &lor{}
    _ixor = &ixor{}
    _lxor = &lxor{}
    _ishl = &ishl{}
    _ishr = &ishr{}
    _iushr = &iushr{}
    _lshl = &lshl{}
    _lshr = &lshr{}
    _lushr = &lushr{}
    _d2f = &d2f{}
    _d2i = &d2i{}
    _d2l = &d2l{}
    _f2d = &f2d{}
    _f2i = &f2i{}
    _f2l = &f2l{}
    _i2b = &i2b{} 
    _i2c = &i2c{}
    _i2d = &i2d{}
    _i2f = &i2f{}
    _i2l = &i2l{}
    _i2s = &i2s{}
    _l2d = &l2d{}
    _l2f = &l2f{}
    _l2i = &l2i{}
    _dcmpg = &dcmpg{}
    _dcmpl = &dcmpl{}
    _fcmpg = &fcmpg{}
    _fcmpl = &fcmpl{}
    _lcmp = &lcmp{}
    _return_ = &return_{}
    _areturn = &areturn{}
    _dreturn = &dreturn{}
    _freturn = &freturn{}
    _ireturn = &ireturn{}
    _lreturn = &lreturn{}
    _monitorenter = &monitorenter{}
    _monitorexit = &monitorexit{}
    _arraylength = &arraylength{}
    _athrow = &athrow{}
)

func Decode(bcr *BytecodeReader) (Instruction) {
    opcode := bcr.readUint8()
    instruction := newInstruction(opcode)
    instruction.fetchOperands(bcr)
    return instruction
}

func newInstruction(opcode byte) (Instruction) {
    switch opcode {
    case 0x00: return _nop
    case 0x01: return _aconst_null
    case 0x02: return _iconst_m1
    case 0x03: return _iconst_0
    case 0x04: return _iconst_1
    case 0x05: return _iconst_2
    case 0x06: return _iconst_3
    case 0x07: return _iconst_4
    case 0x08: return _iconst_5
    case 0x09: return _lconst_0
    case 0x0a: return _lconst_1
    case 0x0b: return _fconst_0
    case 0x0c: return _fconst_1
    case 0x0d: return _fconst_2
    case 0x0e: return _dconst_0
    case 0x0f: return _dconst_1
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
    case 0x1a: return _iload_0
    case 0x1b: return _iload_1
    case 0x1c: return _iload_2
    case 0x1d: return _iload_3
    case 0x1e: return _lload_0
    case 0x1f: return _lload_1
    case 0x20: return _lload_2
    case 0x21: return _lload_3
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
    case 0x2e: return _iaload
    case 0x2f: return _laload
    case 0x30: return _faload
    case 0x31: return _daload
    case 0x32: return _aaload
    case 0x33: return _baload
    case 0x34: return _caload
    case 0x35: return _saload
    case 0x36: return &istore{}
    case 0x37: return &lstore{}
    case 0x38: return &fstore{}
    case 0x39: return &dstore{}
    case 0x3a: return &astore{}
    case 0x3b: return _istore_0
    case 0x3c: return _istore_1
    case 0x3d: return _istore_2
    case 0x3e: return _istore_3
    case 0x3f: return _lstore_0
    case 0x40: return _lstore_1
    case 0x41: return _lstore_2
    case 0x42: return _lstore_3
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
    case 0x4f: return _iastore
    case 0x50: return _lastore
    case 0x51: return _fastore
    case 0x52: return _dastore
    case 0x53: return _aastore
    case 0x54: return _bastore
    case 0x55: return _castore
    case 0x56: return _sastore
    case 0x57: return _pop
    case 0x58: return _pop2
    case 0x59: return _dup
    case 0x5a: return _dup_x1
    case 0x5b: return _dup_x2
    case 0x5c: return _dup2
    case 0x5d: return _dup2_x1
    case 0x5e: return _dup2_x2
    case 0x5f: return _swap
    case 0x60: return _iadd
    case 0x61: return _ladd
    case 0x62: return _fadd
    case 0x63: return _dadd
    case 0x64: return _isub
    case 0x65: return _lsub
    case 0x66: return _fsub
    case 0x67: return _dsub
    case 0x68: return _imul
    case 0x69: return _lmul
    case 0x6a: return _fmul
    case 0x6b: return _dmul
    case 0x6c: return _idiv
    case 0x6d: return _ldiv
    case 0x6e: return _fdiv
    case 0x6f: return _ddiv
    case 0x70: return _irem
    case 0x71: return _lrem
    case 0x72: return _frem
    case 0x73: return _drem
    case 0x74: return _ineg
    case 0x75: return _lneg
    case 0x76: return _fneg
    case 0x77: return _dneg
    case 0x78: return _ishl
    case 0x79: return _lshl
    case 0x7a: return _ishr
    case 0x7b: return _lshr
    case 0x7c: return _iushr
    case 0x7d: return _lushr
    case 0x7e: return _iand
    case 0x7f: return _land
    case 0x80: return _ior
    case 0x81: return _lor
    case 0x82: return _ixor
    case 0x83: return _lxor
    case 0x84: return &iinc{}
    case 0x85: return _i2l
    case 0x86: return _i2f
    case 0x87: return _i2d
    case 0x88: return _l2i
    case 0x89: return _l2f
    case 0x8a: return _l2d
    case 0x8b: return _f2i
    case 0x8c: return _f2l
    case 0x8d: return _f2d
    case 0x8e: return _d2i
    case 0x8f: return _d2l
    case 0x90: return _d2f
    case 0x91: return _i2b
    case 0x92: return _i2c
    case 0x93: return _i2s
    case 0x94: return _lcmp
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
    case 0xa7: return &goto_{}
  //case 0xa8: return &jsr{}
  //case 0xa9: return &ret{}
    case 0xaa: return &tableswitch{}
    case 0xab: return &lookupswitch{}
    case 0xac: return _ireturn
    case 0xad: return _lreturn
    case 0xae: return _freturn
    case 0xaf: return _dreturn
    case 0xb0: return _areturn
    case 0xb1: return _return_
    case 0xb2: return &getstatic{}
    case 0xb3: return &putstatic{}
    case 0xb4: return &getfield{}
    case 0xb5: return &putfield{}
    case 0xb6: return &invokevirtual{}
    case 0xb7: return &invokespecial{}
    case 0xb8: return &invokestatic{}
    case 0xb9: return &invokeinterface{}
    case 0xba: return &invokedynamic{}
    case 0xbb: return &new_{}
    case 0xbc: return &newarray{}
    case 0xbd: return &anewarray{}
    case 0xbe: return _arraylength
    case 0xbf: return _athrow
    case 0xc0: return &checkcast{}
    case 0xc1: return &instanceof{}
    case 0xc2: return _monitorenter
    case 0xc3: return _monitorexit
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
