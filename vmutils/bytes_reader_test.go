package vmutils

import (
	"encoding/binary"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

var data = []byte{
	0x12,
	0x34, 0x56,
	0x78, 0x90, 0xab, 0xcd,
	0xef, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd,
	0x01,
}

func TestBigEndian(t *testing.T) {
	testBigEndian(t, binary.BigEndian)
}

func TestLittleEndian(t *testing.T) {
	testLittleEndian(t, binary.LittleEndian)
}

func TestNativeEndian(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		testLittleEndian(t, NativeEndian)
	}
}

func testBigEndian(t *testing.T, byteOrder binary.ByteOrder) {
	reader := NewBytesReader(data, byteOrder)
	require.Equal(t, uint8(0x12), reader.ReadUint8())
	require.Equal(t, uint16(0x3456), reader.ReadUint16())
	require.Equal(t, uint32(0x7890abcd), reader.ReadUint32())
	require.Equal(t, uint64(0xef1234567890abcd), reader.ReadUint64())
	require.Equal(t, uint8(0x01), reader.ReadUint8())

	reader = NewBytesReader(data, byteOrder)
	require.Equal(t, []byte{0x12, 0x34, 0x56}, reader.ReadBytes(3))
	require.Equal(t, uint16(0x7890), reader.ReadUint16())
	require.Equal(t, uint32(0xabcdef12), reader.ReadUint32())
}

func testLittleEndian(t *testing.T, byteOrder binary.ByteOrder) {
	reader := NewBytesReader(data, byteOrder)
	require.Equal(t, uint8(0x12), reader.ReadUint8())
	require.Equal(t, uint16(0x5634), reader.ReadUint16())
	require.Equal(t, uint32(0xcdab9078), reader.ReadUint32())
	require.Equal(t, uint64(0xcdab9078563412ef), reader.ReadUint64())
	require.Equal(t, uint8(0x01), reader.ReadUint8())

	reader = NewBytesReader(data, byteOrder)
	require.Equal(t, []byte{0x12, 0x34, 0x56}, reader.ReadBytes(3))
	require.Equal(t, uint16(0x9078), reader.ReadUint16())
	require.Equal(t, uint32(0x12efcdab), reader.ReadUint32())
}
