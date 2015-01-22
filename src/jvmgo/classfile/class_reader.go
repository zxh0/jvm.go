package classfile

import (
    "bytes"
    "encoding/binary"
)

type ClassReader struct {
    data    []byte
    //index   uint
    reader  *bytes.Reader
}

func (self *ClassReader) readUint8() (x uint8) {
    read(self, &x)
    return
}

func (self *ClassReader) readUint16() (x uint16) {
    read(self, &x)
    return
}

func (self *ClassReader) readUint32() (x uint32) {
    read(self, &x)
    return
}
func (self *ClassReader) readInt32() (x int32) {
    read(self, &x)
    return
}

func (self *ClassReader) readInt64() (x int64) {
    read(self, &x)
    return
}

func (self *ClassReader) readFloat32() (x float32) {
    read(self, &x)
    return
}
func (self *ClassReader) readFloat64() (x float64) {
    read(self, &x)
    return
}

func read(self *ClassReader, data interface{}) {
    err := binary.Read(self.reader, binary.BigEndian, data)
    if err != nil {
        panic(err.Error())
    }
}

func (self *ClassReader) readBytes(length uint32) ([]byte) {
    bytes := make([]byte, length)
    n, err := self.reader.Read(bytes)
    if err != nil {
        panic(err.Error())
    }
    if n != int(length) {
        // todo
        panic("data not enough!")
    }
    return bytes
}

func (self *ClassReader) readString() (string) {
    length := uint32(self.readUint16())
    bytes := self.readBytes(length)
    return string(bytes[:])
}

// factory
func newClassReader(data []byte) *ClassReader {
    return &ClassReader{data, bytes.NewReader(data)}
}
