package loads

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Load int from local variable
type ILoad struct{ base.Index8Instruction }

func (instr *ILoad) Execute(frame *rtda.Frame) {
	val := frame.LocalVars().GetInt(instr.Index)
	frame.OperandStack().PushInt(val)
}

// Load long from local variable
type LLoad struct{ base.Index8Instruction }

func (instr *LLoad) Execute(frame *rtda.Frame) {
	val := frame.LocalVars().GetLong(instr.Index)
	frame.OperandStack().PushLong(val)
}

// Load float from local variable
type FLoad struct{ base.Index8Instruction }

func (instr *FLoad) Execute(frame *rtda.Frame) {
	val := frame.LocalVars().GetFloat(instr.Index)
	frame.OperandStack().PushFloat(val)
}

// Load double from local variable
type DLoad struct{ base.Index8Instruction }

func (instr *DLoad) Execute(frame *rtda.Frame) {
	val := frame.LocalVars().GetDouble(instr.Index)
	frame.OperandStack().PushDouble(val)
}

// Load reference from local variable
type ALoad struct{ base.Index8Instruction }

func (instr *ALoad) Execute(frame *rtda.Frame) {
	ref := frame.LocalVars().GetRef(instr.Index)
	frame.OperandStack().PushRef(ref)
}
