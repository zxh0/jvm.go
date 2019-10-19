package instructions

import (
	"fmt"

	"github.com/zxh0/jvm.go/instructions/base"
	. "github.com/zxh0/jvm.go/instructions/comparisons"
	. "github.com/zxh0/jvm.go/instructions/constants"
	. "github.com/zxh0/jvm.go/instructions/control"
	. "github.com/zxh0/jvm.go/instructions/conversions"
	. "github.com/zxh0/jvm.go/instructions/extended"
	. "github.com/zxh0/jvm.go/instructions/loads"
	. "github.com/zxh0/jvm.go/instructions/math"
	. "github.com/zxh0/jvm.go/instructions/references"
	. "github.com/zxh0/jvm.go/instructions/reserved"
	. "github.com/zxh0/jvm.go/instructions/stack"
	. "github.com/zxh0/jvm.go/instructions/stores"
)

// NoOperandsInstruction singletons
var (
	nop           = &NOP{}
	aconst_null   = NewConstNull()
	iconst_m1     = NewConstInt(-1)
	iconst_0      = NewConstInt(0)
	iconst_1      = NewConstInt(1)
	iconst_2      = NewConstInt(2)
	iconst_3      = NewConstInt(3)
	iconst_4      = NewConstInt(4)
	iconst_5      = NewConstInt(5)
	lconst_0      = NewConstLong(0)
	lconst_1      = NewConstLong(1)
	fconst_0      = NewConstFloat(0.0)
	fconst_1      = NewConstFloat(1.0)
	fconst_2      = NewConstFloat(2.0)
	dconst_0      = NewConstDouble(0.0)
	dconst_1      = NewConstDouble(1.0)
	iload_0       = NewLoadN(0, false)
	iload_1       = NewLoadN(1, false)
	iload_2       = NewLoadN(2, false)
	iload_3       = NewLoadN(3, false)
	lload_0       = NewLoadN(0, true)
	lload_1       = NewLoadN(1, true)
	lload_2       = NewLoadN(2, true)
	lload_3       = NewLoadN(3, true)
	fload_0       = NewLoadN(0, false)
	fload_1       = NewLoadN(1, false)
	fload_2       = NewLoadN(2, false)
	fload_3       = NewLoadN(3, false)
	dload_0       = NewLoadN(0, true)
	dload_1       = NewLoadN(1, true)
	dload_2       = NewLoadN(2, true)
	dload_3       = NewLoadN(3, true)
	aload_0       = NewLoadN(0, false)
	aload_1       = NewLoadN(1, false)
	aload_2       = NewLoadN(2, false)
	aload_3       = NewLoadN(3, false)
	iaload        = NewIALoad()
	laload        = NewLALoad()
	faload        = NewFALoad()
	daload        = NewDALoad()
	aaload        = NewAALoad()
	baload        = NewBALoad()
	caload        = NewCALoad()
	saload        = NewSALoad()
	istore_0      = NewStoreN(0, false)
	istore_1      = NewStoreN(1, false)
	istore_2      = NewStoreN(2, false)
	istore_3      = NewStoreN(3, false)
	lstore_0      = NewStoreN(0, true)
	lstore_1      = NewStoreN(1, true)
	lstore_2      = NewStoreN(2, true)
	lstore_3      = NewStoreN(3, true)
	fstore_0      = NewStoreN(0, false)
	fstore_1      = NewStoreN(1, false)
	fstore_2      = NewStoreN(2, false)
	fstore_3      = NewStoreN(3, false)
	dstore_0      = NewStoreN(0, true)
	dstore_1      = NewStoreN(1, true)
	dstore_2      = NewStoreN(2, true)
	dstore_3      = NewStoreN(3, true)
	astore_0      = NewStoreN(0, false)
	astore_1      = NewStoreN(1, false)
	astore_2      = NewStoreN(2, false)
	astore_3      = NewStoreN(3, false)
	iastore       = NewIAStore()
	lastore       = NewLAStore()
	fastore       = NewFAStore()
	dastore       = NewDAStore()
	aastore       = NewAAStore()
	bastore       = NewBAStore()
	castore       = NewCAStore()
	sastore       = NewSAStore()
	pop           = &Pop{}
	pop2          = &Pop2{}
	dup           = &Dup{}
	dup_x1        = &DupX1{}
	dup_x2        = &DupX2{}
	dup2          = &Dup2{}
	dup2_x1       = &Dup2X1{}
	dup2_x2       = &Dup2X2{}
	swap          = &Swap{}
	iadd          = NewIAdd()
	ladd          = NewLAdd()
	fadd          = NewFAdd()
	dadd          = NewDAdd()
	isub          = NewISub()
	lsub          = NewLSub()
	fsub          = NewFSub()
	dsub          = NewDSub()
	imul          = NewIMul()
	lmul          = NewLMul()
	fmul          = NewFMul()
	dmul          = NewDMul()
	idiv          = NewIDiv()
	ldiv          = NewLDiv()
	fdiv          = NewFDiv()
	ddiv          = NewDDiv()
	irem          = NewIRem()
	lrem          = NewLRem()
	frem          = NewFRem()
	drem          = NewDRem()
	ineg          = NewINeg()
	lneg          = NewLNeg()
	fneg          = NewFNeg()
	dneg          = NewDNeg()
	ishl          = NewIShl()
	lshl          = NewLShl()
	ishr          = NewIShr()
	lshr          = NewLShr()
	iushr         = NewIUShr()
	lushr         = NewLUShr()
	iand          = NewIAnd()
	land          = NewLAnd()
	ior           = NewIOr()
	lor           = NewLOr()
	ixor          = NewIXor()
	lxor          = NewLXor()
	i2l           = NewI2L()
	i2f           = NewI2F()
	i2d           = NewI2D()
	l2i           = NewL2I()
	l2f           = NewL2F()
	l2d           = NewL2D()
	f2i           = NewF2I()
	f2l           = NewF2L()
	f2d           = NewF2D()
	d2i           = NewD2I()
	d2l           = NewD2L()
	d2f           = NewD2F()
	i2b           = NewI2B()
	i2c           = NewI2C()
	i2s           = NewI2S()
	lcmp          = NewLCMP()
	fcmpl         = NewFCMPL()
	fcmpg         = NewFCMPG()
	dcmpl         = NewDCMPL()
	dcmpg         = NewDCMPG()
	ireturn       = NewXReturn(false)
	lreturn       = NewXReturn(true)
	freturn       = NewXReturn(false)
	dreturn       = NewXReturn(true)
	areturn       = NewXReturn(false)
	_return       = &Return{}
	arraylength   = &ArrayLength{}
	athrow        = &AThrow{}
	monitorenter  = &MonitorEnter{}
	monitorexit   = &MonitorExit{}
	invoke_native = &InvokeNative{}
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
		return &BIPush{}
	case 0x11:
		return &SIPush{}
	case 0x12:
		return &LDC{}
	case 0x13:
		return &LDC_W{}
	case 0x14:
		return &LDC2_W{}
	case 0x15:
		return NewLoad(false)
	case 0x16:
		return NewLoad(true)
	case 0x17:
		return NewLoad(false)
	case 0x18:
		return NewLoad(true)
	case 0x19:
		return NewLoad(false)
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
		return NewStore(false)
	case 0x37:
		return NewStore(true)
	case 0x38:
		return NewStore(false)
	case 0x39:
		return NewStore(true)
	case 0x3a:
		return NewStore(false)
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
		return &IInc{}
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
		return NewIfEQ()
	case 0x9a:
		return NewIfNE()
	case 0x9b:
		return NewIfLT()
	case 0x9c:
		return NewIfGE()
	case 0x9d:
		return NewIfGT()
	case 0x9e:
		return NewIfLE()
	case 0x9f:
		return NewIfICmpEQ()
	case 0xa0:
		return NewIfICmpNE()
	case 0xa1:
		return NewIfICmpLT()
	case 0xa2:
		return NewIfICmpGE()
	case 0xa3:
		return NewIfICmpGT()
	case 0xa4:
		return NewIfICmpLE()
	case 0xa5:
		return NewIfACmpEQ()
	case 0xa6:
		return NewIfACmpNE()
	case 0xa7:
		return &Goto{}
	case 0xa8:
		return &JSR{}
	case 0xa9:
		return &RET{}
	case 0xaa:
		return &TableSwitch{}
	case 0xab:
		return &LookupSwitch{}
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
		return &GetStatic{}
	case 0xb3:
		return &PupStatic{}
	case 0xb4:
		return &GetField{}
	case 0xb5:
		return &PutField{}
	case 0xb6:
		return &InvokeVirtual{}
	case 0xb7:
		return &InvokeSpecial{}
	case 0xb8:
		return &InvokeStatic{}
	case 0xb9:
		return &InvokeInterface{}
	case 0xba:
		return &InvokeDynamic{}
	case 0xbb:
		return &New{}
	case 0xbc:
		return &NewArray{}
	case 0xbd:
		return &ANewArray{}
	case 0xbe:
		return arraylength
	case 0xbf:
		return athrow
	case 0xc0:
		return &CheckCast{}
	case 0xc1:
		return &InstanceOf{}
	case 0xc2:
		return monitorenter
	case 0xc3:
		return monitorexit
	case 0xc4:
		return &Wide{}
	case 0xc5:
		return &MultiANewArray{}
	case 0xc6:
		return NewIfNull()
	case 0xc7:
		return NewIfNonNull()
	case 0xc8:
		return &GotoW{}
	case 0xc9:
		return &JSR_W{}
	//case 0xca: todo breakpoint
	case 0xfe:
		return invoke_native // impdep1
	case 0xff:
		return &Bootstrap{} // impdep2
	default:
		panic(fmt.Errorf("invalid opcode: %v", opcode))
	}
}
