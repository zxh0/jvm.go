package class

import (
    "testing"
    . "jvmgo/test"
)

func TestCalcArgCount(t *testing.T) {
    AssertEquals(0, calcArgCount("()V"))
    AssertEquals(1, calcArgCount("(I)F"))
    AssertEquals(4, calcArgCount("([BIII)V"))
    AssertEquals(3, calcArgCount("(IDLjava/lang/Thread;)Ljava/lang/Object;"))
}
