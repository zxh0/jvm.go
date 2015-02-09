package instructions

import "testing"

func TestReadInt8(t *testing.T) {
    bytecodes := []byte{0xf1, 0x17}
    decoder := &InstructionDecoder{0, bytecodes}
    if b1 := decoder.readInt8(); b1 != -15 {
        t.Errorf("%v", b1)
    }
    if b2 := decoder.readInt8(); b2 != 23 {
        t.Errorf("%v", b2)
    }
}

func TestReadUint8(t *testing.T) {
    bytecodes := []byte{0xf1, 0x17}
    decoder := &InstructionDecoder{0, bytecodes}
    if b1 := decoder.readUint8(); b1 != 241 {
        t.Errorf("%v", b1)
    }
    if b2 := decoder.readUint8(); b2 != 23 {
        t.Errorf("%v", b2)
    }
}

func TestReadInt16(t *testing.T) {
    bytecodes := []byte{0xf1, 0x17}
    decoder := &InstructionDecoder{0, bytecodes}
    if x := decoder.readInt16(); x != -3817 {
        t.Errorf("%v", x)
    }
}

func TestReadUint16(t *testing.T) {
    bytecodes := []byte{0xf1, 0x17}
    decoder := &InstructionDecoder{0, bytecodes}
    if x := decoder.readUint16(); x != 61719 {
        t.Errorf("%v", x)
    }
}
