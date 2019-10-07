package classfile

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAttributeTable(t *testing.T) {
	at := AttributeTable{
		CodeAttribute{MaxStack: 1},
		ConstantValueAttribute{ConstantValueIndex: 2},
		ExceptionsAttribute{ExceptionIndexTable: []uint16{3}},
		BootstrapMethodsAttribute{[]BootstrapMethod{{BootstrapMethodRef: 4}}},
		SignatureAttribute{SignatureIndex: 5},
		SourceFileAttribute{SourceFileIndex: 6},
		LineNumberTableAttribute{[]LineNumberTableEntry{{StartPC: 7}}},
	}
	code, _ := at.GetCodeAttribute()
	require.Equal(t, uint16(1), code.MaxStack)
	require.Equal(t, uint16(2), at.GetConstantValueIndex())
	require.Equal(t, uint16(3), at.GetExceptionIndexTable()[0])
	require.Equal(t, uint16(4), at.GetBootstrapMethods()[0].BootstrapMethodRef)
	require.Equal(t, uint16(5), at.GetSignatureIndex())
	require.Equal(t, uint16(6), at.GetSourceFileIndex())
	require.Equal(t, uint16(7), at.GetLineNumberTable()[0].StartPC)
}
