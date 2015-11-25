package constants

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/jutil"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

// Push item from run-time constant pool
type ldc struct{ base.Index8Instruction }

func (self *ldc) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
}

// Push item from run-time constant pool (wide index)
type ldc_w struct{ base.Index16Instruction }

func (self *ldc_w) Execute(frame *rtda.Frame) {
	_ldc(frame, self.Index)
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
	case *rtc.ConstantClass:
		kClass := c.(*rtc.ConstantClass)
		classObj := kClass.Class().JClass()
		stack.PushRef(classObj)
	default:
		// todo
		// ref to MethodType or MethodHandle
		jutil.Panicf("todo: ldc! %v", c)
	}
}

// Push long or double from run-time constant pool (wide index)
type ldc2_w struct{ base.Index16Instruction }

func (self *ldc2_w) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.ConstantPool()
	c := cp.GetConstant(self.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		jutil.Panicf("ldc2_w! %v", c)
	}
}
