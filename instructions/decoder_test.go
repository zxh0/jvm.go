package instructions

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/zxh0/jvm.go/instructions/base"
	. "github.com/zxh0/jvm.go/instructions/constants"
	. "github.com/zxh0/jvm.go/instructions/loads"
	. "github.com/zxh0/jvm.go/instructions/math"
	. "github.com/zxh0/jvm.go/instructions/stores"
)

func TestDecode(t *testing.T) {
	testConsts(t)
	testLoads(t)
	testStores(t)
	testStack(t)
	testMath(t)
	testConversions(t)
	testComparisons(t)
	testControl(t)
	testReferences(t)
}

func testConsts(t *testing.T) {
	testNoOperands(t, OpAConstNull, aconst_null)
	testNoOperands(t, OpIConstM1, iconst_m1)
	testNoOperands(t, OpIConst0, iconst_0)
	testNoOperands(t, OpIConst1, iconst_1)
	testNoOperands(t, OpIConst2, iconst_2)
	testNoOperands(t, OpIConst3, iconst_3)
	testNoOperands(t, OpIConst4, iconst_4)
	testNoOperands(t, OpIConst5, iconst_5)
	testNoOperands(t, OpLConst0, lconst_0)
	testNoOperands(t, OpLConst1, lconst_1)
	testNoOperands(t, OpFConst0, fconst_0)
	testNoOperands(t, OpFConst1, fconst_1)
	testNoOperands(t, OpFConst2, fconst_2)
	testNoOperands(t, OpDConst0, dconst_0)
	testNoOperands(t, OpDConst1, dconst_1)
	testOperands(t, []byte{OpBIPush, 0x12}, &BIPush{Val: 0x12})
	testOperands(t, []byte{OpSIPush, 0x12, 0x34}, &SIPush{Val: 0x1234})
	testOperands(t, []byte{OpLDC, 0x12}, &LDC{Index8Instruction: Index8Instruction{Index: 0x12}})
	testOperands(t, []byte{OpLDCw, 0x12, 0x34}, &LDC_W{Index16Instruction: Index16Instruction{Index: 0x1234}})
	testOperands(t, []byte{OpLDC2w, 0x12, 0x34}, &LDC2_W{Index16Instruction: Index16Instruction{Index: 0x1234}})
}

func testLoads(t *testing.T) {
	testOperands(t, []byte{OpILoad, 0x12}, newLoad(0x12, false))
	testOperands(t, []byte{OpLLoad, 0x12}, newLoad(0x12, true))
	testOperands(t, []byte{OpFLoad, 0x12}, newLoad(0x12, false))
	testOperands(t, []byte{OpDLoad, 0x12}, newLoad(0x12, true))
	testOperands(t, []byte{OpALoad, 0x12}, newLoad(0x12, false))
	testNoOperands(t, OpILoad0, iload_0)
	testNoOperands(t, OpILoad1, iload_1)
	testNoOperands(t, OpILoad2, iload_2)
	testNoOperands(t, OpILoad3, iload_3)
	testNoOperands(t, OpLLoad0, lload_0)
	testNoOperands(t, OpLLoad1, lload_1)
	testNoOperands(t, OpLLoad2, lload_2)
	testNoOperands(t, OpLLoad3, lload_3)
	testNoOperands(t, OpFLoad0, fload_0)
	testNoOperands(t, OpFLoad1, fload_1)
	testNoOperands(t, OpFLoad2, fload_2)
	testNoOperands(t, OpFLoad3, fload_3)
	testNoOperands(t, OpDLoad0, dload_0)
	testNoOperands(t, OpDLoad1, dload_1)
	testNoOperands(t, OpDLoad2, dload_2)
	testNoOperands(t, OpDLoad3, dload_3)
	testNoOperands(t, OpALoad0, aload_0)
	testNoOperands(t, OpALoad1, aload_1)
	testNoOperands(t, OpALoad2, aload_2)
	testNoOperands(t, OpALoad3, aload_3)
	testNoOperands(t, OpIALoad, iaload)
	testNoOperands(t, OpLALoad, laload)
	testNoOperands(t, OpFALoad, faload)
	testNoOperands(t, OpDALoad, daload)
	testNoOperands(t, OpAALoad, aaload)
	testNoOperands(t, OpBALoad, baload)
	testNoOperands(t, OpCALoad, caload)
	testNoOperands(t, OpSALoad, saload)
}

func newLoad(idx uint, d bool) *Load {
	load := NewLoad(d)
	load.Index = idx
	return load
}

func newStore(idx uint, d bool) *Store {
	store := NewStore(d)
	store.Index = idx
	return store
}

func testStores(t *testing.T) {
	testOperands(t, []byte{OpIStore, 0x12}, newStore(0x12, false))
	testOperands(t, []byte{OpLStore, 0x12}, newStore(0x12, true))
	testOperands(t, []byte{OpFStore, 0x12}, newStore(0x12, false))
	testOperands(t, []byte{OpDStore, 0x12}, newStore(0x12, true))
	testOperands(t, []byte{OpAStore, 0x12}, newStore(0x12, false))
	testNoOperands(t, OpIStore0, istore_0)
	testNoOperands(t, OpIStore1, istore_1)
	testNoOperands(t, OpIStore2, istore_2)
	testNoOperands(t, OpIStore3, istore_3)
	testNoOperands(t, OpLStore0, lstore_0)
	testNoOperands(t, OpLStore1, lstore_1)
	testNoOperands(t, OpLStore2, lstore_2)
	testNoOperands(t, OpLStore3, lstore_3)
	testNoOperands(t, OpFStore0, fstore_0)
	testNoOperands(t, OpFStore1, fstore_1)
	testNoOperands(t, OpFStore2, fstore_2)
	testNoOperands(t, OpFStore3, fstore_3)
	testNoOperands(t, OpDStore0, dstore_0)
	testNoOperands(t, OpDStore1, dstore_1)
	testNoOperands(t, OpDStore2, dstore_2)
	testNoOperands(t, OpDStore3, dstore_3)
	testNoOperands(t, OpAStore0, astore_0)
	testNoOperands(t, OpAStore1, astore_1)
	testNoOperands(t, OpAStore2, astore_2)
	testNoOperands(t, OpAStore3, astore_3)
	testNoOperands(t, OpIAStore, iastore)
	testNoOperands(t, OpLAStore, lastore)
	testNoOperands(t, OpFAStore, fastore)
	testNoOperands(t, OpDAStore, dastore)
	testNoOperands(t, OpAAStore, aastore)
	testNoOperands(t, OpBAStore, bastore)
	testNoOperands(t, OpCAStore, castore)
	testNoOperands(t, OpSAStore, sastore)
}

func testStack(t *testing.T) {
	testNoOperands(t, OpNop, nop)
	testNoOperands(t, OpPop, pop)
	testNoOperands(t, OpPop2, pop2)
	testNoOperands(t, OpDupX1, dup_x1)
	testNoOperands(t, OpDupX2, dup_x2)
	testNoOperands(t, OpDup2, dup2)
	testNoOperands(t, OpDup2X1, dup2_x1)
	testNoOperands(t, OpDup2X2, dup2_x2)
	testNoOperands(t, OpSwap, swap)
}

func testMath(t *testing.T) {
	testNoOperands(t, OpIAdd, iadd)
	testNoOperands(t, OpLAdd, ladd)
	testNoOperands(t, OpFAdd, fadd)
	testNoOperands(t, OpDAdd, dadd)
	testNoOperands(t, OpISub, isub)
	testNoOperands(t, OpLSub, lsub)
	testNoOperands(t, OpFSub, fsub)
	testNoOperands(t, OpDSub, dsub)
	testNoOperands(t, OpIMul, imul)
	testNoOperands(t, OpLMul, lmul)
	testNoOperands(t, OpFMul, fmul)
	testNoOperands(t, OpDMul, dmul)
	testNoOperands(t, OpIDiv, idiv)
	testNoOperands(t, OpLDiv, ldiv)
	testNoOperands(t, OpFDiv, fdiv)
	testNoOperands(t, OpDDiv, ddiv)
	testNoOperands(t, OpIRem, irem)
	testNoOperands(t, OpLRem, lrem)
	testNoOperands(t, OpFRem, frem)
	testNoOperands(t, OpDRem, drem)
	testNoOperands(t, OpINeg, ineg)
	testNoOperands(t, OpLNeg, lneg)
	testNoOperands(t, OpFNeg, fneg)
	testNoOperands(t, OpDNeg, dneg)
	testNoOperands(t, OpIShl, ishl)
	testNoOperands(t, OpLShl, lshl)
	testNoOperands(t, OpIShr, ishr)
	testNoOperands(t, OpLShr, lshr)
	testNoOperands(t, OpIUshr, iushr)
	testNoOperands(t, OpLUshr, lushr)
	testNoOperands(t, OpIAnd, iand)
	testNoOperands(t, OpLAnd, land)
	testNoOperands(t, OpIOr, ior)
	testNoOperands(t, OpLOr, lor)
	testNoOperands(t, OpIXor, ixor)
	testNoOperands(t, OpLXor, lxor)
	testOperands(t, []byte{OpIInc, 0x12, 0x34}, &IInc{Index: 0x12, Const: 0x34})
}

func testConversions(t *testing.T) {
	testNoOperands(t, OpI2L, i2l)
	testNoOperands(t, OpI2F, i2f)
	testNoOperands(t, OpI2D, i2d)
	testNoOperands(t, OpL2I, l2i)
	testNoOperands(t, OpL2F, l2f)
	testNoOperands(t, OpL2D, l2d)
	testNoOperands(t, OpF2I, f2i)
	testNoOperands(t, OpF2L, f2l)
	testNoOperands(t, OpF2D, f2d)
	testNoOperands(t, OpD2I, d2i)
	testNoOperands(t, OpD2L, d2l)
	testNoOperands(t, OpD2F, d2f)
	testNoOperands(t, OpI2B, i2b)
	testNoOperands(t, OpI2C, i2c)
	testNoOperands(t, OpI2S, i2s)
}

func testComparisons(t *testing.T) {
	testNoOperands(t, OpLCmp, lcmp)
	testNoOperands(t, OpFCmpL, fcmpl)
	testNoOperands(t, OpFCmpG, fcmpg)
	testNoOperands(t, OpDCmpL, dcmpl)
	testNoOperands(t, OpDCmpG, dcmpg)
	//testOperands(t, []byte{OpIfEQ, 0x12, 0x34}, &IfEQ{BranchInstruction: BranchInstruction{Offset: 0x1234}})
	//testOperands(t, []byte{OpIfNE, 0x12, 0x34}, &IfNE{BranchInstruction: BranchInstruction{Offset: 0x1234}})
	//testOperands(t, []byte{OpIfLT, 0x12, 0x34}, &IfLT{BranchInstruction: BranchInstruction{Offset: 0x1234}})
	//testOperands(t, []byte{OpIfGE, 0x12, 0x34}, &IfGE{BranchInstruction: BranchInstruction{Offset: 0x1234}})
	//testOperands(t, []byte{OpIfGT, 0x12, 0x34}, &IfGT{BranchInstruction: BranchInstruction{Offset: 0x1234}})
	//testOperands(t, []byte{OpIfLE, 0x12, 0x34}, &IfLE{BranchInstruction: BranchInstruction{Offset: 0x1234}})
	//testOperands(t, []byte{OpIfICmpEQ, 0x12, 0x34}, &IfICmpEQ{BranchInstruction: BranchInstruction{Offset: 0x1234}})
	//testOperands(t, []byte{OpIfICmpNE, 0x12, 0x34}, &IfICmpNE{BranchInstruction: BranchInstruction{Offset: 0x1234}})
	//testOperands(t, []byte{OpIfICmpLT, 0x12, 0x34}, &IfICmpLT{BranchInstruction: BranchInstruction{Offset: 0x1234}})
	//testOperands(t, []byte{OpIfICmpGE, 0x12, 0x34}, &IfICmpGE{BranchInstruction: BranchInstruction{Offset: 0x1234}})
	//testOperands(t, []byte{OpIfICmpGT, 0x12, 0x34}, &IfICmpGT{BranchInstruction: BranchInstruction{Offset: 0x1234}})
	//testOperands(t, []byte{OpIfICmpLE, 0x12, 0x34}, &IfICmpLE{BranchInstruction: BranchInstruction{Offset: 0x1234}})
	//testOperands(t, []byte{OpIfACmpEQ, 0x12, 0x34}, &IfACmpEQ{BranchInstruction: BranchInstruction{Offset: 0x1234}})
	//testOperands(t, []byte{OpIfACmpNE, 0x12, 0x34}, &IfACmpNE{BranchInstruction: BranchInstruction{Offset: 0x1234}})
}

func testControl(t *testing.T) {
	// TODO
	testNoOperands(t, OpIReturn, ireturn)
	testNoOperands(t, OpLReturn, lreturn)
	testNoOperands(t, OpFReturn, freturn)
	testNoOperands(t, OpDReturn, dreturn)
	testNoOperands(t, OpAReturn, areturn)
	testNoOperands(t, OpReturn, _return)
}

func testReferences(t *testing.T) {
	// TODO
	testNoOperands(t, OpArrayLength, arraylength)
	testNoOperands(t, OpAThrow, athrow)
	testNoOperands(t, OpMonitorEnter, monitorenter)
	testNoOperands(t, OpMonitorExit, monitorexit)
	testNoOperands(t, OpInvokeNative, invoke_native)
}

func testNoOperands(t *testing.T, opcode byte, instruction Instruction) {
	testOperands(t, []byte{opcode}, instruction)
}

func testOperands(t *testing.T, code []byte, instruction Instruction) {
	code = append(code, OpNop)
	instructions := Decode(code)
	require.Equal(t, instruction, instructions[0])
	require.Equal(t, nop, instructions[len(code)-1])
}
