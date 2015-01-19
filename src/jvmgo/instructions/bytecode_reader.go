package instructions

type BytecodeReader struct {
    pc          int
    bytecodes   []byte
}

func (self *BytecodeReader) readInt8() (int8) {
    return int8(self.readUint8())
}
func (self *BytecodeReader) readUint8() (uint8) {
    i := self.bytecodes[self.pc]
    self.pc++
    return i
}

func (self *BytecodeReader) readInt16() (int16) {
    return int16(self.readUint16())
}
func (self *BytecodeReader) readUint16() (uint16) {
    byte1 := uint16(self.readUint8())
    byte2 := uint16(self.readUint8())
    return (byte1 << 8) | byte2
}

func (self *BytecodeReader) readInt32() (int32) {
    return int32(self.readUint32())
}
func (self *BytecodeReader) readUint32() (uint32) {
    byte1 := uint32(self.readUint8())
    byte2 := uint32(self.readUint8())
    byte3 := uint32(self.readUint8())
    byte4 := uint32(self.readUint8())
    return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func newBytecodeReader(bytecodes []byte) (*BytecodeReader) {
    return &BytecodeReader{0, bytecodes}
}
