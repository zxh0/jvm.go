package class

import (
    "testing"
    "jvmgo/test"
)

func TestInternSameString(t *testing.T) {
    chars := []uint16{1, 2, 3}
    str := &Obj{}

    str2 := InternString(chars, str)
    test.AssertSame(str, str2)
}
