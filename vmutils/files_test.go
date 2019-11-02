package vmutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStripFileName(t *testing.T) {
	require.Equal(t, "path/to", StripFileName("path/to/Foo.class"))
	require.Equal(t, "path/to", StripFileName("path/to"))
	require.Equal(t, "", StripFileName("Foo.class"))
}
