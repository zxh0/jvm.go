package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	. "github.com/zxh0/jvm.go/jvmgo/instructions/comparisons"
	. "github.com/zxh0/jvm.go/jvmgo/instructions/constants"
	. "github.com/zxh0/jvm.go/jvmgo/instructions/control"
	. "github.com/zxh0/jvm.go/jvmgo/instructions/conversions"
	. "github.com/zxh0/jvm.go/jvmgo/instructions/extended"
	. "github.com/zxh0/jvm.go/jvmgo/instructions/loads"
	. "github.com/zxh0/jvm.go/jvmgo/instructions/math"
	. "github.com/zxh0/jvm.go/jvmgo/instructions/references"
	. "github.com/zxh0/jvm.go/jvmgo/instructions/reserved"
	. "github.com/zxh0/jvm.go/jvmgo/instructions/stack"
	. "github.com/zxh0/jvm.go/jvmgo/instructions/stores"
	"github.com/zxh0/jvm.go/jvmgo/jutil"
)

// NoOperandsInstruction singletons
var (
	_nop           = &NOP{}
	_aconst_null   = &ACONST_NULL{}
	_iconst_m1     = &ICONST_M1{}
	_iconst_0      = &ICONST_0{}
	_iconst_1      = &ICONST_1{}
	_iconst_2      = &ICONST_2{}
	_iconst_3      = &ICONST_3{}
	_iconst_4      = &ICONST_4{}
	_iconst_5      = &ICONST_5{}
	_lconst_0      = &LCONST_0{}
	_lconst_1      = &LCONST_1{}
	_fconst_0      = &FCONST_0{}
	_fconst_1      = &FCONST_1{}
	_fconst_2      = &FCONST_2{}
	_dconst_0      = &DCONST_0{}
	_dconst_1      = &DCONST_1{}
	_iload_0       = &ILOAD_0{}
	_iload_1       = &ILOAD_1{}
	_iload_2       = &ILOAD_2{}
	_iload_3       = &ILOAD_3{}
	_lload_0       = &LLOAD_0{}
	_lload_1       = &LLOAD_1{}
	_lload_2       = &LLOAD_2{}
	_lload_3       = &LLOAD_3{}
	_fload_0       = &FLOAD_0{}
	_fload_1       = &FLOAD_1{}
	_fload_2       = &FLOAD_2{}
	_fload_3       = &FLOAD_3{}
	_dload_0       = &DLOAD_0{}
	_dload_1       = &DLOAD_1{}
	_dload_2       = &DLOAD_2{}
	_dload_3       = &DLOAD_3{}
	_aload_0       = &ALOAD_0{}
	_aload_1       = &ALOAD_1{}
	_aload_2       = &ALOAD_2{}
	_aload_3       = &ALOAD_3{}
	_iaload        = &IALOAD{}
	_laload        = &LALOAD{}
	_faload        = &FALOAD{}
	_daload        = &DALOAD{}
	_aaload        = &AALOAD{}
	_baload        = &BALOAD{}
	_caload        = &CALOAD{}
	_saload        = &SALOAD{}
	_istore_0      = &ISTORE_0{}
	_istore_1      = &ISTORE_1{}
	_istore_2      = &ISTORE_2{}
	_istore_3      = &ISTORE_3{}
	_lstore_0      = &LSTORE_0{}
	_lstore_1      = &LSTORE_1{}
	_lstore_2      = &LSTORE_2{}
	_lstore_3      = &LSTORE_3{}
	_fstore_0      = &FSTORE_0{}
	_fstore_1      = &FSTORE_1{}
	_fstore_2      = &FSTORE_2{}
	_fstore_3      = &FSTORE_3{}
	_dstore_0      = &DSTORE_0{}
	_dstore_1      = &DSTORE_1{}
	_dstore_2      = &DSTORE_2{}
	_dstore_3      = &DSTORE_3{}
	_astore_0      = &ASTORE_0{}
	_astore_1      = &ASTORE_1{}
	_astore_2      = &ASTORE_2{}
	_astore_3      = &ASTORE_3{}
	_iastore       = &IASTORE{}
	_lastore       = &LASTORE{}
	_fastore       = &FASTORE{}
	_dastore       = &DASTORE{}
	_aastore       = &AASTORE{}
	_bastore       = &BASTORE{}
	_castore       = &CASTORE{}
	_sastore       = &SASTORE{}
	_pop           = &POP{}
	_pop2          = &POP2{}
	_dup           = &DUP{}
	_dup_x1        = &DUP_X1{}
	_dup_x2        = &DUP_X2{}
	_dup2          = &DUP2{}
	_dup2_x1       = &DUP2_X1{}
	_dup2_x2       = &DUP2_X2{}
	_swap          = &SWAP{}
	_iadd          = &IADD{}
	_ladd          = &LADD{}
	_fadd          = &FADD{}
	_dadd          = &DADD{}
	_isub          = &ISUB{}
	_lsub          = &LSUB{}
	_fsub          = &FSUB{}
	_dsub          = &DSUB{}
	_imul          = &IMUL{}
	_lmul          = &LMUL{}
	_fmul          = &FMUL{}
	_dmul          = &DMUL{}
	_idiv          = &IDIV{}
	_ldiv          = &LDIV{}
	_fdiv          = &FDIV{}
	_ddiv          = &DDIV{}
	_irem          = &IREM{}
	_lrem          = &LREM{}
	_frem          = &FREM{}
	_drem          = &DREM{}
	_ineg          = &INEG{}
	_lneg          = &LNEG{}
	_fneg          = &FNEG{}
	_dneg          = &DNEG{}
	_ishl          = &ISHL{}
	_lshl          = &LSHL{}
	_ishr          = &ISHR{}
	_lshr          = &LSHR{}
	_iushr         = &IUSHR{}
	_lushr         = &LUSHR{}
	_iand          = &IAND{}
	_land          = &LAND{}
	_ior           = &IOR{}
	_lor           = &LOR{}
	_ixor          = &IXOR{}
	_lxor          = &LXOR{}
	_i2l           = &I2L{}
	_i2f           = &I2F{}
	_i2d           = &I2D{}
	_l2i           = &L2I{}
	_l2f           = &L2F{}
	_l2d           = &L2D{}
	_f2i           = &F2I{}
	_f2l           = &F2L{}
	_f2d           = &F2D{}
	_d2i           = &D2I{}
	_d2l           = &D2L{}
	_d2f           = &D2F{}
	_i2b           = &I2B{}
	_i2c           = &I2C{}
	_i2s           = &I2S{}
	_lcmp          = &LCMP{}
	_fcmpl         = &FCMPL{}
	_fcmpg         = &FCMPG{}
	_dcmpl         = &DCMPL{}
	_dcmpg         = &DCMPG{}
	_ireturn       = &IRETURN{}
	_lreturn       = &LRETURN{}
	_freturn       = &FRETURN{}
	_dreturn       = &DRETURN{}
	_areturn       = &ARETURN{}
	_return_       = &RETURN{}
	_arraylength   = &ARRAY_LENGTH{}
	_athrow        = &ATHROW{}
	_monitorenter  = &MONITOR_ENTER{}
	_monitorexit   = &MONITOR_EXIT{}
	_invoke_native = &INVOKE_NATIVE{}
)

func newInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return _nop
	case 0x01:
		return _aconst_null
	case 0x02:
		return _iconst_m1
	case 0x03:
		return _iconst_0
	case 0x04:
		return _iconst_1
	case 0x05:
		return _iconst_2
	case 0x06:
		return _iconst_3
	case 0x07:
		return _iconst_4
	case 0x08:
		return _iconst_5
	case 0x09:
		return _lconst_0
	case 0x0a:
		return _lconst_1
	case 0x0b:
		return _fconst_0
	case 0x0c:
		return _fconst_1
	case 0x0d:
		return _fconst_2
	case 0x0e:
		return _dconst_0
	case 0x0f:
		return _dconst_1
	case 0x10:
		return &BIPUSH{}
	case 0x11:
		return &SIPUSH{}
	case 0x12:
		return &LDC{}
	case 0x13:
		return &LDC_W{}
	case 0x14:
		return &LDC2_W{}
	case 0x15:
		return &ILOAD{}
	case 0x16:
		return &LLOAD{}
	case 0x17:
		return &FLOAD{}
	case 0x18:
		return &DLOAD{}
	case 0x19:
		return &ALOAD{}
	case 0x1a:
		return _iload_0
	case 0x1b:
		return _iload_1
	case 0x1c:
		return _iload_2
	case 0x1d:
		return _iload_3
	case 0x1e:
		return _lload_0
	case 0x1f:
		return _lload_1
	case 0x20:
		return _lload_2
	case 0x21:
		return _lload_3
	case 0x22:
		return _fload_0
	case 0x23:
		return _fload_1
	case 0x24:
		return _fload_2
	case 0x25:
		return _fload_3
	case 0x26:
		return _dload_0
	case 0x27:
		return _dload_1
	case 0x28:
		return _dload_2
	case 0x29:
		return _dload_3
	case 0x2a:
		return _aload_0
	case 0x2b:
		return _aload_1
	case 0x2c:
		return _aload_2
	case 0x2d:
		return _aload_3
	case 0x2e:
		return _iaload
	case 0x2f:
		return _laload
	case 0x30:
		return _faload
	case 0x31:
		return _daload
	case 0x32:
		return _aaload
	case 0x33:
		return _baload
	case 0x34:
		return _caload
	case 0x35:
		return _saload
	case 0x36:
		return &ISTORE{}
	case 0x37:
		return &LSTORE{}
	case 0x38:
		return &FSTORE{}
	case 0x39:
		return &DSTORE{}
	case 0x3a:
		return &ASTORE{}
	case 0x3b:
		return _istore_0
	case 0x3c:
		return _istore_1
	case 0x3d:
		return _istore_2
	case 0x3e:
		return _istore_3
	case 0x3f:
		return _lstore_0
	case 0x40:
		return _lstore_1
	case 0x41:
		return _lstore_2
	case 0x42:
		return _lstore_3
	case 0x43:
		return _fstore_0
	case 0x44:
		return _fstore_1
	case 0x45:
		return _fstore_2
	case 0x46:
		return _fstore_3
	case 0x47:
		return _dstore_0
	case 0x48:
		return _dstore_1
	case 0x49:
		return _dstore_2
	case 0x4a:
		return _dstore_3
	case 0x4b:
		return _astore_0
	case 0x4c:
		return _astore_1
	case 0x4d:
		return _astore_2
	case 0x4e:
		return _astore_3
	case 0x4f:
		return _iastore
	case 0x50:
		return _lastore
	case 0x51:
		return _fastore
	case 0x52:
		return _dastore
	case 0x53:
		return _aastore
	case 0x54:
		return _bastore
	case 0x55:
		return _castore
	case 0x56:
		return _sastore
	case 0x57:
		return _pop
	case 0x58:
		return _pop2
	case 0x59:
		return _dup
	case 0x5a:
		return _dup_x1
	case 0x5b:
		return _dup_x2
	case 0x5c:
		return _dup2
	case 0x5d:
		return _dup2_x1
	case 0x5e:
		return _dup2_x2
	case 0x5f:
		return _swap
	case 0x60:
		return _iadd
	case 0x61:
		return _ladd
	case 0x62:
		return _fadd
	case 0x63:
		return _dadd
	case 0x64:
		return _isub
	case 0x65:
		return _lsub
	case 0x66:
		return _fsub
	case 0x67:
		return _dsub
	case 0x68:
		return _imul
	case 0x69:
		return _lmul
	case 0x6a:
		return _fmul
	case 0x6b:
		return _dmul
	case 0x6c:
		return _idiv
	case 0x6d:
		return _ldiv
	case 0x6e:
		return _fdiv
	case 0x6f:
		return _ddiv
	case 0x70:
		return _irem
	case 0x71:
		return _lrem
	case 0x72:
		return _frem
	case 0x73:
		return _drem
	case 0x74:
		return _ineg
	case 0x75:
		return _lneg
	case 0x76:
		return _fneg
	case 0x77:
		return _dneg
	case 0x78:
		return _ishl
	case 0x79:
		return _lshl
	case 0x7a:
		return _ishr
	case 0x7b:
		return _lshr
	case 0x7c:
		return _iushr
	case 0x7d:
		return _lushr
	case 0x7e:
		return _iand
	case 0x7f:
		return _land
	case 0x80:
		return _ior
	case 0x81:
		return _lor
	case 0x82:
		return _ixor
	case 0x83:
		return _lxor
	case 0x84:
		return &IINC{}
	case 0x85:
		return _i2l
	case 0x86:
		return _i2f
	case 0x87:
		return _i2d
	case 0x88:
		return _l2i
	case 0x89:
		return _l2f
	case 0x8a:
		return _l2d
	case 0x8b:
		return _f2i
	case 0x8c:
		return _f2l
	case 0x8d:
		return _f2d
	case 0x8e:
		return _d2i
	case 0x8f:
		return _d2l
	case 0x90:
		return _d2f
	case 0x91:
		return _i2b
	case 0x92:
		return _i2c
	case 0x93:
		return _i2s
	case 0x94:
		return _lcmp
	case 0x95:
		return _fcmpl
	case 0x96:
		return _fcmpg
	case 0x97:
		return _dcmpl
	case 0x98:
		return _dcmpg
	case 0x99:
		return &IFEQ{}
	case 0x9a:
		return &IFNE{}
	case 0x9b:
		return &IFLT{}
	case 0x9c:
		return &IFGE{}
	case 0x9d:
		return &IFGT{}
	case 0x9e:
		return &IFLE{}
	case 0x9f:
		return &IF_ICMPEQ{}
	case 0xa0:
		return &IF_ICMPNE{}
	case 0xa1:
		return &IF_ICMPLT{}
	case 0xa2:
		return &IF_ICMPGE{}
	case 0xa3:
		return &IF_ICMPGT{}
	case 0xa4:
		return &IF_ICMPLE{}
	case 0xa5:
		return &IF_ACMPEQ{}
	case 0xa6:
		return &IF_ACMPNE{}
	case 0xa7:
		return &GOTO{}
	case 0xa8:
		return &JSR{}
	case 0xa9:
		return &RET{}
	case 0xaa:
		return &TABLE_SWITCH{}
	case 0xab:
		return &LOOKUP_SWITCH{}
	case 0xac:
		return _ireturn
	case 0xad:
		return _lreturn
	case 0xae:
		return _freturn
	case 0xaf:
		return _dreturn
	case 0xb0:
		return _areturn
	case 0xb1:
		return _return_
	case 0xb2:
		return &GET_STATIC{}
	case 0xb3:
		return &PUT_STATIC{}
	case 0xb4:
		return &GET_FIELD{}
	case 0xb5:
		return &PUT_FIELD{}
	case 0xb6:
		return &INVOKE_VIRTUAL{}
	case 0xb7:
		return &INVOKE_SPECIAL{}
	case 0xb8:
		return &INVOKE_STATIC{}
	case 0xb9:
		return &INVOKE_INTERFACE{}
	case 0xba:
		return &INVOKE_DYNAMIC{}
	case 0xbb:
		return &NEW{}
	case 0xbc:
		return &NEW_ARRAY{}
	case 0xbd:
		return &ANEW_ARRAY{}
	case 0xbe:
		return _arraylength
	case 0xbf:
		return _athrow
	case 0xc0:
		return &CHECK_CAST{}
	case 0xc1:
		return &INSTANCE_OF{}
	case 0xc2:
		return _monitorenter
	case 0xc3:
		return _monitorexit
	case 0xc4:
		return &WIDE{}
	case 0xc5:
		return &MULTI_ANEW_ARRAY{}
	case 0xc6:
		return &IFNULL{}
	case 0xc7:
		return &IFNONNULL{}
	case 0xc8:
		return &GOTO_W{}
	case 0xc9:
		return &JSR_W{}
	//case 0xca: todo breakpoint
	case 0xfe:
		return _invoke_native // impdep1
	case 0xff:
		return &BOOTSTRAP{} // impdep2
	default:
		jutil.Panicf("BAD opcode: %v!", opcode)
		return nil
	}
}
