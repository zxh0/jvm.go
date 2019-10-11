package rtda

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zxh0/jvm.go/rtda/heap"
)

func TestLocalVars(t *testing.T) {
	lv := newLocalVars(16)
	obj := &heap.Object{}

	lv.SetIntVar(0, 123)
	lv.SetLongVar(1, 456)
	lv.SetFloatVar(2, 3.14)
	lv.SetDoubleVar(3, 2.71828)
	lv.SetRefVar(4, obj)

	require.Equal(t, int32(123), lv.GetIntVar(0))
	require.Equal(t, int64(456), lv.GetLongVar(1))
	require.Equal(t, float32(3.14), lv.GetFloatVar(2))
	require.Equal(t, 2.71828, lv.GetDoubleVar(3))
	require.Equal(t, obj, lv.GetRefVar(4))
}

func TestOperandStack(t *testing.T) {
	obj := &heap.Object{}

	stack := newOperandStack(16)
	require.True(t, stack.IsStackEmpty())

	stack.PushInt(123)
	stack.PushLong(456)
	stack.PushFloat(3.14)
	stack.PushDouble(2.71828)
	stack.PushRef(obj)

	require.False(t, stack.IsStackEmpty())
	require.Equal(t, obj, stack.PopRef())
	require.Equal(t, 2.71828, stack.PopDouble())
	require.Equal(t, float32(3.14), stack.PopFloat())
	require.Equal(t, int64(456), stack.PopLong())
	require.Equal(t, int32(123), stack.PopInt())
	require.True(t, stack.IsStackEmpty())
}
