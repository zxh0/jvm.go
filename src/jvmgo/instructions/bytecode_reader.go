package instructions

type BytecodeReader struct {
    pc          int
    bytecodes   []byte
}

func newBytecodeReader(bytecodes []byte) (*BytecodeReader) {
    return &BytecodeReader{0, bytecodes}
}
