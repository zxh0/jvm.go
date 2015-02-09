package instructions

import (
    "testing"
    . "jvmgo/test"
)

func TestReadInt8(t *testing.T) {
    bytecodes := []byte{0xf1, 0x17}
    decoder := &InstructionDecoder{0, bytecodes}
    AssertEquals(-15, decoder.readInt8())
    AssertEquals(23, decoder.readInt8())
}

func TestReadUint8(t *testing.T) {
    bytecodes := []byte{0xf1, 0x17}
    decoder := &InstructionDecoder{0, bytecodes}
    AssertEquals(241, decoder.readUint8())
    AssertEquals(23, decoder.readUint8())
}

func TestReadInt16(t *testing.T) {
    bytecodes := []byte{0xf1, 0x17}
    decoder := &InstructionDecoder{0, bytecodes}
    AssertEquals(-3817, decoder.readInt16())
}

func TestReadUint16(t *testing.T) {
    bytecodes := []byte{0xf1, 0x17}
    decoder := &InstructionDecoder{0, bytecodes}
    AssertEquals(61719, decoder.readUint16())
}
