package rtda

import (
    "testing"
    //"jvmgo/rtda/class"
)

func TestPushPop(t *testing.T) {
    maxSize := 20
    stack := newStack(maxSize)
    
    for i := 0; i < maxSize; i++ {
        stack.push(newFrame(i))
    }
    //stack.push(newFrame(0))

    for i := maxSize - 1; i >= 0 ; i-- {
        frame := stack.pop()
        if frame.nextPC != i {
            t.Errorf("i %v", i)
        }
    }
}

func TestIsEmpty(t *testing.T) {
    stack := newStack(17)
    if empty := stack.isEmpty(); !empty {
        t.Errorf("not empty!")
    }

    stack.push(newFrame(1))
    if empty := stack.isEmpty(); empty {
        t.Errorf("empty!")
    }
}

func newFrame(nextPC int) (*Frame) {
    frame := &Frame{}
    frame.nextPC = nextPC
    return frame
}
