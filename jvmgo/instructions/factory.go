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
	nop           = &NOP{}
	aconst_null   = &ACONST_NULL{}
	iconst_m1     = &ICONST_M1{}
	iconst_0      = &ICONST_0{}
	iconst_1      = &ICONST_1{}
	iconst_2      = &ICONST_2{}
	iconst_3      = &ICONST_3{}
	iconst_4      = &ICONST_4{}
	iconst_5      = &ICONST_5{}
	lconst_0      = &LCONST_0{}
	lconst_1      = &LCONST_1{}
	fconst_0      = &FCONST_0{}
	fconst_1      = &FCONST_1{}
	fconst_2      = &FCONST_2{}
	dconst_0      = &DCONST_0{}
	dconst_1      = &DCONST_1{}
	iload_0       = &ILOAD_0{}
	iload_1       = &ILOAD_1{}
	iload_2       = &ILOAD_2{}
	iload_3       = &ILOAD_3{}
	lload_0       = &LLOAD_0{}
	lload_1       = &LLOAD_1{}
	lload_2       = &LLOAD_2{}
	lload_3       = &LLOAD_3{}
	fload_0       = &FLOAD_0{}
	fload_1       = &FLOAD_1{}
	fload_2       = &FLOAD_2{}
	fload_3       = &FLOAD_3{}
	dload_0       = &DLOAD_0{}
	dload_1       = &DLOAD_1{}
	dload_2       = &DLOAD_2{}
	dload_3       = &DLOAD_3{}
	aload_0       = &ALOAD_0{}
	aload_1       = &ALOAD_1{}
	aload_2       = &ALOAD_2{}
	aload_3       = &ALOAD_3{}
	iaload        = &IALOAD{}
	laload        = &LALOAD{}
	faload        = &FALOAD{}
	daload        = &DALOAD{}
	aaload        = &AALOAD{}
	baload        = &BALOAD{}
	caload        = &CALOAD{}
	saload        = &SALOAD{}
	istore_0      = &ISTORE_0{}
	istore_1      = &ISTORE_1{}
	istore_2      = &ISTORE_2{}
	istore_3      = &ISTORE_3{}
	lstore_0      = &LSTORE_0{}
	lstore_1      = &LSTORE_1{}
	lstore_2      = &LSTORE_2{}
	lstore_3      = &LSTORE_3{}
	fstore_0      = &FSTORE_0{}
	fstore_1      = &FSTORE_1{}
	fstore_2      = &FSTORE_2{}
	fstore_3      = &FSTORE_3{}
	dstore_0      = &DSTORE_0{}
	dstore_1      = &DSTORE_1{}
	dstore_2      = &DSTORE_2{}
	dstore_3      = &DSTORE_3{}
	astore_0      = &ASTORE_0{}
	astore_1      = &ASTORE_1{}
	astore_2      = &ASTORE_2{}
	astore_3      = &ASTORE_3{}
	iastore       = &IASTORE{}
	lastore       = &LASTORE{}
	fastore       = &FASTORE{}
	dastore       = &DASTORE{}
	aastore       = &AASTORE{}
	bastore       = &BASTORE{}
	castore       = &CASTORE{}
	sastore       = &SASTORE{}
	pop           = &POP{}
	pop2          = &POP2{}
	dup           = &DUP{}
	dup_x1        = &DUP_X1{}
	dup_x2        = &DUP_X2{}
	dup2          = &DUP2{}
	dup2_x1       = &DUP2_X1{}
	dup2_x2       = &DUP2_X2{}
	swap          = &SWAP{}
	iadd          = &IADD{}
	ladd          = &LADD{}
	fadd          = &FADD{}
	dadd          = &DADD{}
	isub          = &ISUB{}
	lsub          = &LSUB{}
	fsub          = &FSUB{}
	dsub          = &DSUB{}
	imul          = &IMUL{}
	lmul          = &LMUL{}
	fmul          = &FMUL{}
	dmul          = &DMUL{}
	idiv          = &IDIV{}
	ldiv          = &LDIV{}
	fdiv          = &FDIV{}
	ddiv          = &DDIV{}
	irem          = &IREM{}
	lrem          = &LREM{}
	frem          = &FREM{}
	drem          = &DREM{}
	ineg          = &INEG{}
	lneg          = &LNEG{}
	fneg          = &FNEG{}
	dneg          = &DNEG{}
	ishl          = &ISHL{}
	lshl          = &LSHL{}
	ishr          = &ISHR{}
	lshr          = &LSHR{}
	iushr         = &IUSHR{}
	lushr         = &LUSHR{}
	iand          = &IAND{}
	land          = &LAND{}
	ior           = &IOR{}
	lor           = &LOR{}
	ixor          = &IXOR{}
	lxor          = &LXOR{}
	i2l           = &I2L{}
	i2f           = &I2F{}
	i2d           = &I2D{}
	l2i           = &L2I{}
	l2f           = &L2F{}
	l2d           = &L2D{}
	f2i           = &F2I{}
	f2l           = &F2L{}
	f2d           = &F2D{}
	d2i           = &D2I{}
	d2l           = &D2L{}
	d2f           = &D2F{}
	i2b           = &I2B{}
	i2c           = &I2C{}
	i2s           = &I2S{}
	lcmp          = &LCMP{}
	fcmpl         = &FCMPL{}
	fcmpg         = &FCMPG{}
	dcmpl         = &DCMPL{}
	dcmpg         = &DCMPG{}
	ireturn       = &IRETURN{}
	lreturn       = &LRETURN{}
	freturn       = &FRETURN{}
	dreturn       = &DRETURN{}
	areturn       = &ARETURN{}
	_return       = &RETURN{}
	arraylength   = &ARRAY_LENGTH{}
	athrow        = &ATHROW{}
	monitorenter  = &MONITOR_ENTER{}
	monitorexit   = &MONITOR_EXIT{}
	invoke_native = &INVOKE_NATIVE{}
)

func newInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	case 0x02:
		return iconst_m1
	case 0x03:
		return iconst_0
	case 0x04:
		return iconst_1
	case 0x05:
		return iconst_2
	case 0x06:
		return iconst_3
	case 0x07:
		return iconst_4
	case 0x08:
		return iconst_5
	case 0x09:
		return lconst_0
	case 0x0a:
		return lconst_1
	case 0x0b:
		return fconst_0
	case 0x0c:
		return fconst_1
	case 0x0d:
		return fconst_2
	case 0x0e:
		return dconst_0
	case 0x0f:
		return dconst_1
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
		return iload_0
	case 0x1b:
		return iload_1
	case 0x1c:
		return iload_2
	case 0x1d:
		return iload_3
	case 0x1e:
		return lload_0
	case 0x1f:
		return lload_1
	case 0x20:
		return lload_2
	case 0x21:
		return lload_3
	case 0x22:
		return fload_0
	case 0x23:
		return fload_1
	case 0x24:
		return fload_2
	case 0x25:
		return fload_3
	case 0x26:
		return dload_0
	case 0x27:
		return dload_1
	case 0x28:
		return dload_2
	case 0x29:
		return dload_3
	case 0x2a:
		return aload_0
	case 0x2b:
		return aload_1
	case 0x2c:
		return aload_2
	case 0x2d:
		return aload_3
	case 0x2e:
		return iaload
	case 0x2f:
		return laload
	case 0x30:
		return faload
	case 0x31:
		return daload
	case 0x32:
		return aaload
	case 0x33:
		return baload
	case 0x34:
		return caload
	case 0x35:
		return saload
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
		return istore_0
	case 0x3c:
		return istore_1
	case 0x3d:
		return istore_2
	case 0x3e:
		return istore_3
	case 0x3f:
		return lstore_0
	case 0x40:
		return lstore_1
	case 0x41:
		return lstore_2
	case 0x42:
		return lstore_3
	case 0x43:
		return fstore_0
	case 0x44:
		return fstore_1
	case 0x45:
		return fstore_2
	case 0x46:
		return fstore_3
	case 0x47:
		return dstore_0
	case 0x48:
		return dstore_1
	case 0x49:
		return dstore_2
	case 0x4a:
		return dstore_3
	case 0x4b:
		return astore_0
	case 0x4c:
		return astore_1
	case 0x4d:
		return astore_2
	case 0x4e:
		return astore_3
	case 0x4f:
		return iastore
	case 0x50:
		return lastore
	case 0x51:
		return fastore
	case 0x52:
		return dastore
	case 0x53:
		return aastore
	case 0x54:
		return bastore
	case 0x55:
		return castore
	case 0x56:
		return sastore
	case 0x57:
		return pop
	case 0x58:
		return pop2
	case 0x59:
		return dup
	case 0x5a:
		return dup_x1
	case 0x5b:
		return dup_x2
	case 0x5c:
		return dup2
	case 0x5d:
		return dup2_x1
	case 0x5e:
		return dup2_x2
	case 0x5f:
		return swap
	case 0x60:
		return iadd
	case 0x61:
		return ladd
	case 0x62:
		return fadd
	case 0x63:
		return dadd
	case 0x64:
		return isub
	case 0x65:
		return lsub
	case 0x66:
		return fsub
	case 0x67:
		return dsub
	case 0x68:
		return imul
	case 0x69:
		return lmul
	case 0x6a:
		return fmul
	case 0x6b:
		return dmul
	case 0x6c:
		return idiv
	case 0x6d:
		return ldiv
	case 0x6e:
		return fdiv
	case 0x6f:
		return ddiv
	case 0x70:
		return irem
	case 0x71:
		return lrem
	case 0x72:
		return frem
	case 0x73:
		return drem
	case 0x74:
		return ineg
	case 0x75:
		return lneg
	case 0x76:
		return fneg
	case 0x77:
		return dneg
	case 0x78:
		return ishl
	case 0x79:
		return lshl
	case 0x7a:
		return ishr
	case 0x7b:
		return lshr
	case 0x7c:
		return iushr
	case 0x7d:
		return lushr
	case 0x7e:
		return iand
	case 0x7f:
		return land
	case 0x80:
		return ior
	case 0x81:
		return lor
	case 0x82:
		return ixor
	case 0x83:
		return lxor
	case 0x84:
		return &IINC{}
	case 0x85:
		return i2l
	case 0x86:
		return i2f
	case 0x87:
		return i2d
	case 0x88:
		return l2i
	case 0x89:
		return l2f
	case 0x8a:
		return l2d
	case 0x8b:
		return f2i
	case 0x8c:
		return f2l
	case 0x8d:
		return f2d
	case 0x8e:
		return d2i
	case 0x8f:
		return d2l
	case 0x90:
		return d2f
	case 0x91:
		return i2b
	case 0x92:
		return i2c
	case 0x93:
		return i2s
	case 0x94:
		return lcmp
	case 0x95:
		return fcmpl
	case 0x96:
		return fcmpg
	case 0x97:
		return dcmpl
	case 0x98:
		return dcmpg
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
		return ireturn
	case 0xad:
		return lreturn
	case 0xae:
		return freturn
	case 0xaf:
		return dreturn
	case 0xb0:
		return areturn
	case 0xb1:
		return _return
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
		return arraylength
	case 0xbf:
		return athrow
	case 0xc0:
		return &CHECK_CAST{}
	case 0xc1:
		return &INSTANCE_OF{}
	case 0xc2:
		return monitorenter
	case 0xc3:
		return monitorexit
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
		return invoke_native // impdep1
	case 0xff:
		return &BOOTSTRAP{} // impdep2
	default:
		jutil.Panicf("BAD opcode: %v!", opcode)
		return nil
	}
}
