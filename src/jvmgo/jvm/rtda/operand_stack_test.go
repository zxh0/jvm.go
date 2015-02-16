package rtda

import (
	rtc "jvmgo/jvm/rtda/class"
	. "jvmgo/test"
	"testing"
)

func TestPushAndPop(t *testing.T) {
	stack := newOperandStack(6)
	stack.PushNull()
	stack.PushRef(&rtc.Obj{})
	stack.PushInt(-37)
	stack.PushLong(0xabcd1234ff)
	stack.PushFloat(3.14)
	stack.PushDouble(-2.71828)
	//stack.PushInt(0)

	AssertEquals(-2.71828, stack.PopDouble())
	AssertEquals(3.14, stack.PopFloat())
	AssertEquals(int64(0xabcd1234ff), stack.PopLong())
	AssertEquals(-37, stack.PopInt())
	AssertNotNil(stack.PopRef())
	AssertNil(stack.PopRef())
}

func TestPopN(t *testing.T) {
	stack := newOperandStack(6)
	stack.PushInt(4)
	stack.PushInt(5)
	stack.PushInt(6)
	stack.PushInt(8)
	stack.PushInt(9)

	top3 := stack.popN(3)
	AssertEquals([]int32{6, 8, 9}, top3)
	AssertEquals(5, stack.PopInt())
}

func TestTop(t *testing.T) {
	stack := newOperandStack(3)
	stack.PushInt(1)
	stack.PushInt(2)
	stack.PushInt(3)

	AssertEquals(3, stack.Top(0))
	AssertEquals(2, stack.Top(1))
	AssertEquals(1, stack.Top(2))
}
