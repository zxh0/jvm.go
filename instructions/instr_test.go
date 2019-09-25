package instructions

import (
	"github.com/zxh0/jvm.go/rtda/heap"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

func TestAConst(t *testing.T) {
	frame := rtda.NewFrame(0, 1)
	aconst_null.Execute(frame)
}

func TestDConst(t *testing.T) {
	frame := rtda.NewFrame(0, 1)
	dconst_0.Execute(frame)
	require.Equal(t, 0.0, frame.OperandStack().PopDouble())
	dconst_1.Execute(frame)
	require.Equal(t, 1.0, frame.OperandStack().PopDouble())
}

func TestFConst(t *testing.T) {
	frame := rtda.NewFrame(0, 1)
	fconst_0.Execute(frame)
	require.Equal(t, float32(0.0), frame.OperandStack().PopFloat())
	fconst_1.Execute(frame)
	require.Equal(t, float32(1.0), frame.OperandStack().PopFloat())
	fconst_2.Execute(frame)
	require.Equal(t, float32(2.0), frame.OperandStack().PopFloat())
}

func TestIConst(t *testing.T) {
	frame := rtda.NewFrame(0, 1)
	iconst_m1.Execute(frame)
	require.Equal(t, int32(-1), frame.OperandStack().PopInt())
	iconst_0.Execute(frame)
	require.Equal(t, int32(0), frame.OperandStack().PopInt())
	iconst_1.Execute(frame)
	require.Equal(t, int32(1), frame.OperandStack().PopInt())
	iconst_2.Execute(frame)
	require.Equal(t, int32(2), frame.OperandStack().PopInt())
	iconst_3.Execute(frame)
	require.Equal(t, int32(3), frame.OperandStack().PopInt())
	iconst_4.Execute(frame)
	require.Equal(t, int32(4), frame.OperandStack().PopInt())
	iconst_5.Execute(frame)
	require.Equal(t, int32(5), frame.OperandStack().PopInt())
}

func TestLConst(t *testing.T) {
	frame := rtda.NewFrame(0, 1)
	lconst_0.Execute(frame)
	require.Equal(t, int64(0), frame.OperandStack().PopLong())
	lconst_1.Execute(frame)
	require.Equal(t, int64(1), frame.OperandStack().PopLong())
}

func TestILoad(t *testing.T) {
	frame := rtda.NewFrame(8, 1)
	instrs := []base.Instruction{iload_0, iload_1, iload_2, iload_3}
	for i, instr := range instrs {
		frame.LocalVars().SetInt(uint(i), 100)
		instr.Execute(frame)
		require.Equal(t, int32(100), frame.OperandStack().PopInt())
	}
}

func TestLLoad(t *testing.T) {
	frame := rtda.NewFrame(8, 1)
	instrs := []base.Instruction{lload_0, lload_1, lload_2, lload_3}
	for i, instr := range instrs {
		frame.LocalVars().SetLong(uint(i), 100)
		instr.Execute(frame)
		require.Equal(t, int64(100), frame.OperandStack().PopLong())
	}
}

func TestFLoad(t *testing.T) {
	frame := rtda.NewFrame(8, 1)
	instrs := []base.Instruction{fload_0, fload_1, fload_2, fload_3}
	for i, instr := range instrs {
		frame.LocalVars().SetFloat(uint(i), 100)
		instr.Execute(frame)
		require.Equal(t, float32(100), frame.OperandStack().PopFloat())
	}
}

func TestDLoad(t *testing.T) {
	frame := rtda.NewFrame(8, 1)
	instrs := []base.Instruction{dload_0, dload_1, dload_2, dload_3}
	for i, instr := range instrs {
		frame.LocalVars().SetDouble(uint(i), 100)
		instr.Execute(frame)
		require.Equal(t, float64(100), frame.OperandStack().PopDouble())
	}
}

func TestALoad(t *testing.T) {
	frame := rtda.NewFrame(8, 1)
	instrs := []base.Instruction{aload_0, aload_1, aload_2, aload_3}
	for i, instr := range instrs {
		obj := &heap.Object{}
		frame.LocalVars().SetRef(uint(i), obj)
		instr.Execute(frame)
		require.Equal(t, obj, frame.OperandStack().PopRef())
	}
}

func TestIStore(t *testing.T) {
	frame := rtda.NewFrame(8, 1)
	instrs := []base.Instruction{istore_0, istore_1, istore_2, istore_3}
	for i, instr := range instrs {
		frame.OperandStack().PushInt(100)
		instr.Execute(frame)
		require.Equal(t, int32(100), frame.LocalVars().GetInt(uint(i)))
	}
}

func TestLStore(t *testing.T) {
	frame := rtda.NewFrame(8, 1)
	instrs := []base.Instruction{lstore_0, lstore_1, lstore_2, lstore_3}
	for i, instr := range instrs {
		frame.OperandStack().PushLong(100)
		instr.Execute(frame)
		require.Equal(t, int64(100), frame.LocalVars().GetLong(uint(i)))
	}
}

func TestFStore(t *testing.T) {
	frame := rtda.NewFrame(8, 1)
	instrs := []base.Instruction{fstore_0, fstore_1, fstore_2, fstore_3}
	for i, instr := range instrs {
		frame.OperandStack().PushFloat(100)
		instr.Execute(frame)
		require.Equal(t, float32(100), frame.LocalVars().GetFloat(uint(i)))
	}
}

func TestDStore(t *testing.T) {
	frame := rtda.NewFrame(8, 1)
	instrs := []base.Instruction{dstore_0, dstore_1, dstore_2, dstore_3}
	for i, instr := range instrs {
		frame.OperandStack().PushDouble(100)
		instr.Execute(frame)
		require.Equal(t, float64(100), frame.LocalVars().GetDouble(uint(i)))
	}
}

func TestAStore(t *testing.T) {
	frame := rtda.NewFrame(8, 1)
	instrs := []base.Instruction{astore_0, astore_1, astore_2, astore_3}
	for i, instr := range instrs {
		obj := &heap.Object{}
		frame.OperandStack().PushRef(obj)
		instr.Execute(frame)
		require.Equal(t, obj, frame.LocalVars().GetRef(uint(i)))
	}
}
