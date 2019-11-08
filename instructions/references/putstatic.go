package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/rtda/linker"
)

// Set static field in class
type PupStatic struct {
	base.Index16Instruction
	field *heap.Field
}

func (instr *PupStatic) Execute(frame *rtda.Frame) {
	if instr.field == nil {
		cp := frame.GetConstantPool()
		fieldRef := cp.GetConstantFieldRef(instr.Index)
		instr.field = linker.ResolveField(frame.GetBootLoader(), fieldRef, true)
	}

	class := instr.field.Class
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread.InitClass(class)
		return
	}

	val := frame.PopL(instr.field.IsLongOrDouble)
	instr.field.PutStaticValue(val)
}
