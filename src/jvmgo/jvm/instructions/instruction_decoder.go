package instructions

type BytecodeReader struct {
    pc      int
    code    []byte // bytecodes
}

func NewBytecodeReader() (*BytecodeReader) {
    return &BytecodeReader{}
}

// todo
func (self *BytecodeReader) Decode() (uint8, Instruction) {
    opcode := self.readUint8()
    instruction := newInstruction(opcode)
    instruction.fetchOperands(self)
    return opcode, instruction
}

// getters & setters
func (self *BytecodeReader) PC() (int) {
    return self.pc
}
func (self *BytecodeReader) SetPC(pc int) {
    self.pc = pc
}
func (self *BytecodeReader) SetCode(code []byte) {
    self.code = code
}

func (self *BytecodeReader) readInt8() (int8) {
    return int8(self.readUint8())
}
func (self *BytecodeReader) readUint8() (uint8) {
    i := self.code[self.pc]
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

func (self *BytecodeReader) readInt32s(count int32) ([]int32) {
    ints := make([]int32, count)
    for i := range ints {
        ints[i] = self.readInt32()
    }
    return ints
}
