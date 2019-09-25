package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Get static field from class
type GET_STATIC struct {
	base.Index16Instruction
	field *heap.Field
}

func (instr *GET_STATIC) Execute(frame *rtda.Frame) {
	if instr.field == nil {
		cp := frame.Method().Class().ConstantPool()
		kFieldRef := cp.GetConstant(instr.Index).(*heap.ConstantFieldref)
		instr.field = kFieldRef.StaticField()
	}

	class := instr.field.Class()
	if class.InitializationNotStarted() {
		frame.RevertNextPC() // undo getstatic
		frame.Thread().InitClass(class)
		return
	}

	val := instr.field.GetStaticValue()
	stack := frame.OperandStack()
	stack.PushField(val, instr.field.IsLongOrDouble)
}
