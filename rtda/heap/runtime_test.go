package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStr(t *testing.T) {
	value, coder := goStrToJSFields("hello")
	require.Equal(t, []int8{'h', 'e', 'l', 'l', 'o'}, value)
	require.Equal(t, int32(0), coder)

	value, coder = goStrToJSFields("你好")
	require.Equal(t, []int8{0x60, 0x4f, 0x7d, 0x59}, value)
	require.Equal(t, int32(1), coder)
}
