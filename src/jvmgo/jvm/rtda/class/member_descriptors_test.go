package class

import (
    "testing"
    . "jvmgo/test"
)

func TestCalcArgCount(t *testing.T) {
    AssertEquals(0, parseMethodDescriptor("()V").argCount())
    AssertEquals(1, parseMethodDescriptor("(I)F").argCount())
    AssertEquals(4, parseMethodDescriptor("([BIII)V").argCount())
    AssertEquals(3, parseMethodDescriptor("(IDLjava/lang/Thread;)Ljava/lang/Object;").argCount())
}
