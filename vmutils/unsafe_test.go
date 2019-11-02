package vmutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCasts(t *testing.T) {
	bytes := []byte("hello, world")
	int8s := CastBytesToInt8s(bytes)
	uint16s := CastInt8sToUint16s(int8s)
	uint32s := CastBytesToUint32s(bytes)
	int32s := CastBytesToInt32s(bytes)

	require.Equal(t, 12, len(int8s))
	require.Equal(t, 6, len(uint16s))
	require.Equal(t, 3, len(uint32s))
	require.Equal(t, 3, len(int32s))

	require.Equal(t, bytes, CastInt8sToBytes(int8s))
	require.Equal(t, int8s, CastUint16sToInt8s(uint16s))
}
