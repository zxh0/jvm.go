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
    byte1 := self.readUint8()
    byte2 := self.readUint8()
    return (uint16(byte1) << 8) | uint16(byte2)
}

func newBytecodeReader(bytecodes []byte) (*BytecodeReader) {
    return &BytecodeReader{0, bytecodes}
}
