package instructions

import (
	"jvmgo/jvm/rtda"
	rtc "jvmgo/jvm/rtda/class"
)

// Invoke instance method; dispatch based on class
type invokevirtual struct{ Index16Instruction }

func (self *invokevirtual) Execute(frame *rtda.Frame) {
	cp := frame.Method().ConstantPool()
	kMethodRef := cp.GetConstant(self.index).(*rtc.ConstantMethodref)

	stack := frame.OperandStack()
	ref := stack.Top(kMethodRef.ArgCount())
	if ref == nil {
		panic("NPE")
	}

	method := kMethodRef.GetVirtualMethod(ref.(*rtc.Obj))
	frame.Thread().InvokeMethod(method)
}
