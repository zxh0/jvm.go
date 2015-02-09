package classfile

import (
    "jvmgo/util/bigendian"
)

type ClassReader struct {
    data []byte
}

func newClassReader(data []byte) *ClassReader {
    return &ClassReader{data}
}

func (self *ClassReader) readUint8() uint8 {
    val := self.data[0]
    self.data = self.data[1:]
    return val
}

func (self *ClassReader) readUint16() uint16 {
    val := bigendian.Uint16(self.data)
    self.data = self.data[2:]
    return val
}

func (self *ClassReader) readUint32() uint32 {
    val := bigendian.Int32(self.data)
    self.data = self.data[4:]
    return uint32(val)
}
func (self *ClassReader) readInt32() int32 {
    val := bigendian.Int32(self.data)
    self.data = self.data[4:]
    return val
}

func (self *ClassReader) readInt64() int64 {
    val := bigendian.Int64(self.data)
    self.data = self.data[8:]
    return val
}

func (self *ClassReader) readFloat32() float32 {
    val := bigendian.Float32(self.data)
    self.data = self.data[4:]
    return val
}

func (self *ClassReader) readFloat64() float64 {
    val := bigendian.Float64(self.data)
    self.data = self.data[8:]
    return val
}

func (self *ClassReader) readBytes(length uint32) ([]byte) {
    bytes := self.data[:length]
    self.data = self.data[length:]
    return bytes
}

// todo
func (self *ClassReader) readString() (string) {
    length := uint32(self.readUint16())
    bytes := self.readBytes(length)
    return string(bytes[:])
}
