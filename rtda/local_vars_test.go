package rtda

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zxh0/jvm.go/rtda/heap"
)

func TestLocalVars(t *testing.T) {
	lv := newLocalVars(16)
	obj := &heap.Object{}

	lv.SetInt(0, 123)
	lv.SetLong(1, 456)
	lv.SetFloat(2, 3.14)
	lv.SetDouble(3, 2.71828)
	lv.SetRef(4, obj)

	require.Equal(t, int32(123), lv.GetInt(0))
	require.Equal(t, int64(456), lv.GetLong(1))
	require.Equal(t, float32(3.14), lv.GetFloat(2))
	require.Equal(t, 2.71828, lv.GetDouble(3))
	require.Equal(t, obj, lv.GetRef(4))
}
