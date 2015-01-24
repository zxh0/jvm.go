package classfile

import (
    "testing"
)

func TestCalcArgCount(t *testing.T) {
    if count := calcArgCount("(IDLjava/lang/Thread;)Ljava/lang/Object;"); count != 3 {
        t.Errorf("%v", count)
    }
}
