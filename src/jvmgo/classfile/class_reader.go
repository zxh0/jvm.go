package classfile

import (
    "bytes"
    "encoding/binary"
)

type ClassReader struct {
    //data    []byte
    //index   uint
    reader  *bytes.Reader
}

// todo
func (self *ClassReader) readUint8() (x uint8) {
    readVal(self, &x)
    return
}

func (self *ClassReader) readUint16() (x uint16) {
    readVal(self, &x)
    return
}

func (self *ClassReader) readUint32() (x uint32) {
    readVal(self, &x)
    return
}
func (self *ClassReader) readInt32() (x int32) {
    readVal(self, &x)
    return
}

func (self *ClassReader) readInt64() (x int64) {
    readVal(self, &x)
    return
}

func (self *ClassReader) readFloat32() (x float32) {
    readVal(self, &x)
    return
}
func (self *ClassReader) readFloat64() (x float64) {
    readVal(self, &x)
    return
}

func readVal(self *ClassReader, data interface{}) {
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

func (self *ClassReader) unreadUint8() {
    err := self.reader.UnreadByte()
    if err != nil {
       panic(err.Error()) 
    }
}

// factory
func newClassReader(data []byte) *ClassReader {
    return &ClassReader{bytes.NewReader(data)}
}
