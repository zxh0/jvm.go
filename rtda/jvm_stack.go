package rtda

import (
	"fmt"
)

// jvm stack
type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame // stack is implemented as linked list
}

func newStack(maxSize uint) *Stack {
	return &Stack{maxSize, 0, nil}
}

func (stack *Stack) isEmpty() bool {
	return stack._top == nil
}

func (stack *Stack) push(frame *Frame) {
	if stack.size >= stack.maxSize {
		// todo
		panic("StackOverflowError")
	}

	if stack._top != nil {
		frame.lower = stack._top
	}

	stack._top = frame
	stack.size++
}

func (stack *Stack) pop() *Frame {
	if stack._top == nil {
		panic("jvm stack is empty!")
	}

	top := stack._top
	stack._top = top.lower
	top.lower = nil
	stack.size--

	return top
}

func (stack *Stack) clear() {
	for !stack.isEmpty() {
		stack.pop()
	}
}

func (stack *Stack) top() *Frame {
	if stack._top == nil {
		panic("jvm stack is empty!")
	}

	return stack._top
}

func (stack *Stack) topN(n uint) *Frame {
	if stack.size < n {
		panic(fmt.Sprintf("jvm stack size:%v n:%v", stack.size, n))
	}

	frame := stack._top
	for n > 0 {
		frame = frame.lower
		n--
	}

	return frame
}
