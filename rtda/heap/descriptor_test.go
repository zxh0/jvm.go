package heap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseMethodDescriptor(t *testing.T) {
	testParseMethodDescriptor(t, "(IDLjava/lang/Thread;)Ljava/lang/Object;")
	testParseMethodDescriptor(t, "([Ljava/lang/Thread;)I")
	testParseMethodDescriptor(t, "()V")
}

func testParseMethodDescriptor(t *testing.T, md string) {
	parsed := parseMethodDescriptor(md)
	require.Equal(t, md, mdToStr(parsed))
}

func mdToStr(md MethodDescriptor) string {
	str := "("
	for _, paramType := range md.ParameterTypes {
		str += string(paramType)
	}
	str += ")"
	str += string(md.ReturnType)
	return str
}
