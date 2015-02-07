package util

import (
    "testing"
    . "jvmgo/test"
)

func TestCastInt8sToUint8s(t *testing.T) {
    a := []int8{-1, 1, 0, -2, 2}
    b := CastInt8sToUint8s(a)
    AssertEquals(0xFF,  b[0])
    AssertEquals(1,     b[1])
    AssertEquals(0,     b[2])
    AssertEquals(0xFE,  b[3])
    AssertEquals(2,     b[4])
}
