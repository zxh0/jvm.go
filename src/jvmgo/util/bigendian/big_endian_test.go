package bigendian

import (
    "testing"
    . "jvmgo/test"
)

func TestBigEndian(t *testing.T) {
    s := make([]byte, 100)

    EncodeInt16(s, 500)
    AssertEquals(500, DecodeInt16(s))
    EncodeInt16(s, -500)
    AssertEquals(-500, DecodeInt16(s))

    EncodeInt32(s, 65539)
    AssertEquals(65539, DecodeInt32(s))
    EncodeInt32(s, -65539)
    AssertEquals(-65539, DecodeInt32(s))

    EncodeInt64(s, 71)
    AssertEquals(71, DecodeInt64(s))
    EncodeInt64(s, -71)
    AssertEquals(-71, DecodeInt64(s))

    EncodeFloat32(s, 3.14)
    AssertEquals(3.14, DecodeFloat32(s))
    EncodeFloat32(s, -3.14)
    AssertEquals(-3.14, DecodeFloat32(s))

    EncodeFloat64(s, 2.71828)
    AssertEquals(2.71828, DecodeFloat64(s))
    EncodeFloat64(s, -2.71828)
    AssertEquals(-2.71828, DecodeFloat64(s))
}
