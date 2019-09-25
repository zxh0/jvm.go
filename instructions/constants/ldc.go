package constants

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/jutil"
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
	stack := frame.OperandStack()
	cp := frame.ConstantPool()
	c := cp.GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string:
		internedStr := rtda.JString(c.(string))
		stack.PushRef(internedStr)
	case *heap.ConstantClass:
		kClass := c.(*heap.ConstantClass)
		classObj := kClass.Class().JClass()
		stack.PushRef(classObj)
	default:
		// todo
		// ref to MethodType or MethodHandle
		jutil.Panicf("todo: ldc! %v", c)
	}
}

// Push long or double from run-time constant pool (wide index)
type LDC2_W struct{ base.Index16Instruction }

func (instr *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.ConstantPool()
	c := cp.GetConstant(instr.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		jutil.Panicf("ldc2_w! %v", c)
	}
}
