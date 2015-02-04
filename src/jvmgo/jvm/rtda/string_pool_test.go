package rtda

import (
    "testing"
    "jvmgo/test"
)

func TestInternSameString(t *testing.T) {
    chars := []uint16{1, 1, 1}
    str := &Obj{}

    str2 := InternString(chars, str)
    test.AssertSame(str, str2)

    str3 := InternString(chars, str2)
    test.AssertSame(str2, str3)
}

func TestInternDifferentStrings(t *testing.T) {
    str1, chars1 := &Obj{extra:1}, []uint16{3, 2, 1}
    str2, chars2 := &Obj{extra:2}, []uint16{1, 2, 3, 4}
    str3, chars3 := &Obj{extra:3}, []uint16{1, 2, 3}

    test.AssertSame(str1, InternString(chars1, str1))
    test.AssertSame(str2, InternString(chars2, str2))
    test.AssertSame(str3, InternString(chars3, str3))

    test.AssertSame(str1, InternString(chars1, &Obj{}))
    test.AssertSame(str2, InternString(chars2, &Obj{}))
    test.AssertSame(str3, InternString(chars3, &Obj{}))
}
