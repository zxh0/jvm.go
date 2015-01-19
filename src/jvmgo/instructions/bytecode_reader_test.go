package instructions

import "testing"

func TestReadInt8(t *testing.T) {
    bytecodes := []byte{0xf1, 0x17}
    bcr := newBytecodeReader(bytecodes)
    if b1 := bcr.readInt8(); b1 != -15 {
        t.Errorf("%v", b1)
    }
    if b2 := bcr.readInt8(); b2 != 23 {
        t.Errorf("%v", b2)
    }
}

func TestReadUint8(t *testing.T) {
    bytecodes := []byte{0xf1, 0x17}
    bcr := newBytecodeReader(bytecodes)
    if b1 := bcr.readUint8(); b1 != 241 {
        t.Errorf("%v", b1)
    }
    if b2 := bcr.readUint8(); b2 != 23 {
        t.Errorf("%v", b2)
    }
}
