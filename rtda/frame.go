package rtda

import (
	"github.com/zxh0/jvm.go/rtda/heap"
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

// TODO
func NewFrame(maxLocals, maxStack int) *Frame {
	return &Frame{
		localVars:    newLocalVars(uint(maxLocals)),
		operandStack: newOperandStack(uint(maxStack)),
	}
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

func (frame *Frame) reset(method *heap.Method) {
	frame.method = method
	frame.nextPC = 0
	frame.lower = nil
	frame.onPopAction = nil
	if frame.maxLocals > 0 {
		frame.localVars.clear()
	}
	if frame.maxStack > 0 {
		frame.operandStack.Clear()
	}
}

// getters & setters
func (frame *Frame) Thread() *Thread {
	return frame.thread
}
func (frame *Frame) Method() *heap.Method {
	return frame.method
}
func (frame *Frame) LocalVars() *LocalVars {
	return frame.localVars
}
func (frame *Frame) OperandStack() *OperandStack {
	return frame.operandStack
}
func (frame *Frame) NextPC() int {
	return frame.nextPC
}
func (frame *Frame) SetNextPC(nextPC int) {
	frame.nextPC = nextPC
}
func (frame *Frame) SetOnPopAction(f func()) {
	frame.onPopAction = f
}

func (frame *Frame) RevertNextPC() {
	frame.nextPC = frame.thread.pc
}

// todo
func (frame *Frame) ClassLoader() *heap.ClassLoader {
	return heap.BootLoader()
}
func (frame *Frame) ConstantPool() *heap.ConstantPool {
	return frame.method.ConstantPool()
}

func (frame *Frame) Load(idx uint, isLongOrDouble bool) {
	slot := frame.LocalVars().Get(idx)
	frame.OperandStack().PushSlot(slot)
	if isLongOrDouble {
		frame.OperandStack().PushNull()
	}
}
func (frame *Frame) Store(idx uint, isLongOrDouble bool) {
	if isLongOrDouble {
		frame.OperandStack().PopSlot()
	}
	slot := frame.OperandStack().PopSlot()
	frame.LocalVars().Set(idx, slot)
}
