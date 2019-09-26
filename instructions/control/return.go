package control

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Return void from method
type Return struct{ base.NoOperandsInstruction }

func (instr *Return) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	thread.PopFrame()
}

// Return reference from method
type AReturn struct{ base.NoOperandsInstruction }

func (instr *AReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.PopRef()
	invokerFrame.PushRef(ref)
}

// Return double from method
type DReturn struct{ base.NoOperandsInstruction }

func (instr *DReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.PopDouble()
	invokerFrame.PushDouble(val)
}

// Return float from method
type FReturn struct{ base.NoOperandsInstruction }

func (instr *FReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.PopFloat()
	invokerFrame.PushFloat(val)
}

// Return int from method
type IReturn struct{ base.NoOperandsInstruction }

func (instr *IReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.PopInt()
	invokerFrame.PushInt(val)
}

// Return double from method
type LReturn struct{ base.NoOperandsInstruction }

func (instr *LReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.PopLong()
	invokerFrame.PushLong(val)
}
