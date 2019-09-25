package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Set static field in class
type PupStatic struct {
	base.Index16Instruction
	field *heap.Field
}

func (instr *PupStatic) Execute(frame *rtda.Frame) {
	if instr.field == nil {
		cp := frame.Method().Class().ConstantPool()
		kFieldRef := cp.GetConstant(instr.Index).(*heap.ConstantFieldref)
		instr.field = kFieldRef.StaticField()
	}

	class := instr.field.Class()
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread().InitClass(class)
		return
	}

	val := frame.OperandStack().PopField(instr.field.IsLongOrDouble)
	instr.field.PutStaticValue(val)
}
