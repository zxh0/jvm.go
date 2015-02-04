package rtda

import (
    "testing"
    "jvmgo/test"
)

func TestPushPop(t *testing.T) {
    maxSize := 20
    stack := newStack(uint(maxSize))
    
    for i := 0; i < maxSize; i++ {
        stack.push(_newFrame(i))
    }
    //stack.push(_newFrame(0))

    for i := maxSize - 1; i >= 0 ; i-- {
        frame := stack.pop()
        test.AssertEquals(i, frame.nextPC)
    }
}

func TestIsEmpty(t *testing.T) {
    stack := newStack(17)
    test.AssertTrue(stack.isEmpty())

    stack.push(_newFrame(1))
    test.AssertFalse(stack.isEmpty())
}

func _newFrame(nextPC int) (*Frame) {
    frame := &Frame{}
    frame.nextPC = nextPC
    return frame
}
