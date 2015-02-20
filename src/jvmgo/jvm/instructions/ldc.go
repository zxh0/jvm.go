package instructions

import (
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
	"jvmgo/util"
)

// Push item from run-time constant pool
type ldc struct{ Index8Instruction }

func (self *ldc) Execute(frame *rtda.Frame) {
	_ldc(frame, self.index)
}

// Push item from run-time constant pool (wide index)
type ldc_w struct{ Index16Instruction }

func (self *ldc_w) Execute(frame *rtda.Frame) {
	_ldc(frame, self.index)
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case *rtc.ConstantString:
		kString := c.(*rtc.ConstantString)
		if kString.JStr() == nil {
			strObj := rtda.NewJString(kString.GoStr(), frame) // already interned
			kString.SetJStr(strObj)
		}
		stack.PushRef(kString.JStr())
	case *rtc.ConstantClass: // todo
		kClass := c.(*rtc.ConstantClass)
		classObj := kClass.Class().JClass()
		stack.PushRef(classObj)
	default:
		// todo
		// ref to MethodType or MethodHandle
		util.Panicf("todo: ldc! %v\n", c)
	}
}

// Push long or double from run-time constant pool (wide index)
type ldc2_w struct{ Index16Instruction }

func (self *ldc2_w) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(self.index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
		// todo
	}
}
