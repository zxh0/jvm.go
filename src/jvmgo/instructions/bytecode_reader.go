package instructions

type BytecodeReader struct {
    pc          int
    bytecodes   []byte
}

func (self *BytecodeReader) readUint8() (uint8) {
    i := self.bytecodes[self.pc]
    self.pc++
    return i
}

func (self *BytecodeReader) readInt8() (int8) {
    return int8(self.readUint8())
}

func newBytecodeReader(bytecodes []byte) (*BytecodeReader) {
    return &BytecodeReader{0, bytecodes}
}
