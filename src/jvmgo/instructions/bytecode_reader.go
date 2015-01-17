package instructions

type BytecodeReader struct {
    pc          int
    bytecodes   []byte
}

func (self *BytecodeReader) readUint8() (byte) {
    b := self.bytecodes[self.pc]
    self.pc++
    return b
}

func newBytecodeReader(bytecodes []byte) (*BytecodeReader) {
    return &BytecodeReader{0, bytecodes}
}
