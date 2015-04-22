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

func (self *Stack) isEmpty() bool {
	return self._top == nil
}

func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		// todo
		panic("StackOverflowError")
	}

	if self._top != nil {
		frame.lower = self._top
	}

	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}

	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--

	return top
}

func (self *Stack) clear() {
	for !self.isEmpty() {
		self.pop()
	}
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}

	return self._top
}

func (self *Stack) topN(n uint) *Frame {
	if self.size < n {
		panic(fmt.Sprintf("jvm stack size:%v n:%v", self.size, n))
	}

	frame := self._top
	for n > 0 {
		frame = frame.lower
		n--
	}

	return frame
}
