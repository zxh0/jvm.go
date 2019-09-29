package rtda

import (
	"github.com/zxh0/jvm.go/rtda/heap"
)

// stack frame
type Frame struct {
	LocalVars
	OperandStack
	lower       *Frame // stack is implemented as linked list
	thread      *Thread
	method      *heap.Method
	maxLocals   uint
	maxStack    uint
	nextPC      int // the next instruction after the call
	onPopAction func()
}

// TODO
func NewFrame(maxLocals, maxStack int) *Frame {
	return &Frame{
		LocalVars:    newLocalVars(uint(maxLocals)),
		OperandStack: newOperandStack(uint(maxStack)),
	}
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		maxLocals:    method.MaxLocals,
		maxStack:     method.MaxStack,
		LocalVars:    newLocalVars(method.MaxLocals),
		OperandStack: newOperandStack(method.MaxStack),
	}
}

func (frame *Frame) reset(method *heap.Method) {
	frame.method = method
	frame.nextPC = 0
	frame.lower = nil
	frame.onPopAction = nil
	if frame.maxLocals > 0 {
		frame.clearLocalVars()
	}
	if frame.maxStack > 0 {
		frame.ClearStack()
	}
}

// getters & setters
func (frame *Frame) Thread() *Thread {
	return frame.thread
}
func (frame *Frame) Method() *heap.Method {
	return frame.method
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

func (frame *Frame) Load(idx uint, isLongOrDouble bool) {
	slot := frame.GetLocalVar(idx)
	frame.Push(slot)
	if isLongOrDouble {
		frame.PushNull()
	}
}
func (frame *Frame) Store(idx uint, isLongOrDouble bool) {
	if isLongOrDouble {
		frame.Pop()
	}
	slot := frame.Pop()
	frame.SetLocalVar(idx, slot)
}

// shortcuts
func (frame *Frame) GetClass() *heap.Class {
	return frame.method.Class
}
func (frame *Frame) GetConstantPool() *heap.ConstantPool {
	return frame.method.Class.ConstantPool
}
