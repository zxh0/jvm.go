package classfile

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadTable(t *testing.T) {
	readFn := func(reader *ClassReader) uint16 {
		return reader.ReadUint16()
	}

	reader := newClassReader([]byte{0x00, 0x00})
	s := reader.readTable(readFn)
	require.Equal(t, []uint16{}, s)

	reader = newClassReader([]byte{0x00, 0x01, 0x02, 0x03})
	s = reader.readTable(readFn)
	require.Equal(t, []uint16{0x0203}, s)
}
