package classfile

import (
    "testing"
)

func TestCalcArgCount(t *testing.T) {
    if count := calcArgCount("()V"); count != 0 {
        t.Errorf("%v", count)
    }
    if count := calcArgCount("(I)F"); count != 1 {
        t.Errorf("%v", count)
    }
    if count := calcArgCount("([BIII])V"); count != 3 {
        t.Errorf("%v", count)
    }
    if count := calcArgCount("(IDLjava/lang/Thread;)Ljava/lang/Object;"); count != 3 {
        t.Errorf("%v", count)
    }
}
