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

// readFn: func(reader *ClassReader) XXX
func (reader *ClassReader) readTable(readFn interface{}) interface{} {
	n := int(reader.ReadUint16())

	itemType := reflect.TypeOf(readFn).Out(0)
	sliceType := reflect.SliceOf(itemType)
	s := reflect.MakeSlice(sliceType, n, n) // make([]x, n, n)

	readFnVal := reflect.ValueOf(readFn)
	args := []reflect.Value{reflect.ValueOf(reader)}

	for i := 0; i < n; i++ {
		x := readFnVal.Call(args)[0]
		s.Index(i).Set(x) // s[i] = x
	}

	return s.Interface()
}
