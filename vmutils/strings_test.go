package vmutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeMUTF8(t *testing.T) {
	require.Equal(t, "\u0000", DecodeMUTF8([]byte{0xc0, 0x80}))
	require.Equal(t, "ß", DecodeMUTF8([]byte{0xc3, 0x9f}))
	require.Equal(t, "abc", DecodeMUTF8([]byte("abc")))
	require.Equal(t, "foo\u0000bar", DecodeMUTF8([]byte("foo\u0000bar"))) // ?
	require.Equal(t, "foo\u0000bar", DecodeMUTF8([]byte("foo\xc0\x80bar")))
	//require.Equal(t, "foo\U00010000bar", DecodeMUTF8([]byte("foo\uD800\uDC00bar")))
}
