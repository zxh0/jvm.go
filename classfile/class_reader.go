package classfile

import (
	"encoding/binary"
	"reflect"

	"github.com/zxh0/jvm.go/vmutils"
)

type ClassReader struct {
	vmutils.BytesReader
	cf *ClassFile
}

func newClassReader(data []byte) ClassReader {
	br := vmutils.NewBytesReader(data, binary.BigEndian)
	return ClassReader{BytesReader: br}
}

func (reader *ClassReader) readUint16s() []uint16 {
	n := reader.ReadUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = reader.ReadUint16()
	}
	return s
}

func (reader *ClassReader) readTable(x interface{},
	readFn func(reader *ClassReader) interface{}) interface{} {

	n := int(reader.ReadUint16())
	s := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(x)), n, n) // make([]x, n, n)

	for i := 0; i < n; i++ {
		x := readFn(reader)
		s.Index(i).Set(reflect.ValueOf(x)) // s[i] = x
	}

	return s.Interface()
}
