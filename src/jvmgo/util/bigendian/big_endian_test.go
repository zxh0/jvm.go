package bigendian

import (
    "testing"
    . "jvmgo/test"
)

func TestBigEndian(t *testing.T) {
    s := make([]byte, 100)

    PutInt16(s, 500)
    AssertEquals(500, Int16(s))
    PutInt16(s, -500)
    AssertEquals(-500, Int16(s))

    PutInt32(s, 65539)
    AssertEquals(65539, Int32(s))
    PutInt32(s, -65539)
    AssertEquals(-65539, Int32(s))

    PutInt64(s, 71)
    AssertEquals(71, Int64(s))
    PutInt64(s, -71)
    AssertEquals(-71, Int64(s))

    PutFloat32(s, 3.14)
    AssertEquals(3.14, Float32(s))
    PutFloat32(s, -3.14)
    AssertEquals(-3.14, Float32(s))

    PutFloat64(s, 2.71828)
    AssertEquals(2.71828, Float64(s))
    PutFloat64(s, -2.71828)
    AssertEquals(-2.71828, Float64(s))
}
