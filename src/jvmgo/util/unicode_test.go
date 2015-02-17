package util

import (
	. "jvmgo/test"
	"testing"
)

func TestParseCommandFail0(t *testing.T) {
	str := "abcd1234@@汉字&中国"
	utf16 := StringToUtf16(str)
	utf8 := Utf16ToString(utf16)
	AssertSame(str, utf8)
}
