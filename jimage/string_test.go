package jimage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmaskedHashCode(t *testing.T) {
	require.Equal(t, int32(16777619), unmaskedHashCode("", HashMultiplier))
	require.Equal(t, int32(1213053849), unmaskedHashCode("foo", HashMultiplier))
	require.Equal(t, int32(977475810), unmaskedHashCode("bar", HashMultiplier))
	require.Equal(t, int32(-1678740824), unmaskedHashCode("Hello, World!", HashMultiplier))
	require.Equal(t, int32(1641313752), unmaskedHashCode("你好，世界！", HashMultiplier))
	require.Equal(t, int32(-348596783), unmaskedHashCode("123456789:一二三四五六七八九", HashMultiplier))
}
