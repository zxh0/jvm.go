package constants

import (
	"fmt"

	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Push item from run-time constant pool
type LDC struct{ base.Index8Instruction }

func (instr *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, instr.Index)
}

// Push item from run-time constant pool (wide index)
type LDC_W struct{ base.Index16Instruction }

func (instr *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, instr.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	c := frame.GetConstantPool().GetConstant(index)

	switch x := c.(type) {
	case int32:
		frame.PushInt(x)
	case float32:
		frame.PushFloat(x)
	case *heap.ConstantString: // string
		frame.PushRef(x.GetJString())
	case *heap.ConstantClass:
		frame.PushRef(x.GetClass().JClass)
	default:
		// todo
		// ref to MethodType or MethodHandle
		panic(fmt.Errorf("todo: ldc! %v", c))
	}
}

// Push long or double from run-time constant pool (wide index)
type LDC2_W struct{ base.Index16Instruction }

func (instr *LDC2_W) Execute(frame *rtda.Frame) {
	c := frame.GetConstantPool().GetConstant(instr.Index)

	switch c.(type) {
	case int64:
		frame.PushLong(c.(int64))
	case float64:
		frame.PushDouble(c.(float64))
	default:
		panic(fmt.Errorf("ldc2_w! %v", c))
	}
}
