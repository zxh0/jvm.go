package rtda

import (
	"github.com/zxh0/jvm.go/rtda/heap"
)

// stack frame
type Frame struct {
	LocalVars
	OperandStack
	lower       *Frame // stack is implemented as linked list
	Thread      *Thread
	Method      *heap.Method
	maxLocals   uint
	maxStack    uint
	NextPC      int // the next instruction after the call
	OnPopAction func()
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
		Thread:       thread,
		Method:       method,
		maxLocals:    method.MaxLocals,
		maxStack:     method.MaxStack,
		LocalVars:    newLocalVars(method.MaxLocals),
		OperandStack: newOperandStack(method.MaxStack),
	}
}

func (frame *Frame) reset(method *heap.Method) {
	frame.Method = method
	frame.NextPC = 0
	frame.lower = nil
	frame.OnPopAction = nil
	if frame.maxLocals > 0 {
		frame.clearLocalVars()
	}
	if frame.maxStack > 0 {
		frame.ClearStack()
	}
}

func (frame *Frame) RevertNextPC() {
	frame.NextPC = frame.Thread.pc
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
	return frame.Method.Class
}
func (frame *Frame) GetConstantPool() heap.ConstantPool {
	return frame.Method.Class.ConstantPool
}

// todo
func (frame *Frame) GetClassLoader() *heap.ClassLoader {
	return heap.BootLoader()
}
