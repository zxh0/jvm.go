package rtda

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

// stack frame
type Frame struct {
	lower        *Frame // stack is implemented as linked list
	thread       *Thread
	method       *heap.Method
	localVars    *LocalVars
	operandStack *OperandStack
	maxLocals    uint
	maxStack     uint
	nextPC       int // the next instruction after the call
	onPopAction  func()
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		maxLocals:    method.MaxLocals(),
		maxStack:     method.MaxStack(),
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

func (self *Frame) reset(method *heap.Method) {
	self.method = method
	self.nextPC = 0
	self.lower = nil
	self.onPopAction = nil
	if self.maxLocals > 0 {
		self.localVars.clear()
	}
	if self.maxStack > 0 {
		self.operandStack.Clear()
	}
}

// getters & setters
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) Method() *heap.Method {
	return self.method
}
func (self *Frame) LocalVars() *LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
func (self *Frame) SetOnPopAction(f func()) {
	self.onPopAction = f
}

func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}

// todo
func (self *Frame) ClassLoader() *heap.ClassLoader {
	return heap.BootLoader()
}
func (self *Frame) ConstantPool() *heap.ConstantPool {
	return self.method.ConstantPool()
}
