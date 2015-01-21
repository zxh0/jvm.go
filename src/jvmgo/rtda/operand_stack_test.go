package rtda

import (
    "testing"
    "jvmgo/rtda/class"
)

func TestPushAndPop(t *testing.T) {
    stack := newOperandStack(8)
    stack.PushNull()
    stack.PushRef(class.NewArray(4, 4))

    if x := stack.PopRef(); x == nil {
        t.Errorf("nil!")
    }
    if x := stack.PopRef(); x != nil {
        t.Errorf("not nil: %v", x)
    }
}
