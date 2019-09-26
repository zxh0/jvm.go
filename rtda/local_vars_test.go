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
