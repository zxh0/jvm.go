package rtda

import (
	. "jvmgo/testing"
	"testing"
)

func TestPushPop(t *testing.T) {
	maxSize := 20
	stack := newStack(uint(maxSize))

	for i := 0; i < maxSize; i++ {
		stack.push(_newFrame(i))
	}
	//stack.push(_newFrame(0))

	for i := maxSize - 1; i >= 0; i-- {
		AssertEquals(i, stack.pop().nextPC)
	}
}

func TestTopN(t *testing.T) {
	stack := newStack(10)
	stack.push(_newFrame(5))
	stack.push(_newFrame(6))
	stack.push(_newFrame(7))

	AssertEquals(7, stack.topN(0).nextPC)
	AssertEquals(6, stack.topN(1).nextPC)
	AssertEquals(5, stack.topN(2).nextPC)
}

func TestIsEmpty(t *testing.T) {
	stack := newStack(17)
	AssertTrue(stack.isEmpty())

	stack.push(_newFrame(1))
	AssertFalse(stack.isEmpty())
}

func _newFrame(nextPC int) *Frame {
	frame := &Frame{}
	frame.nextPC = nextPC
	return frame
}
