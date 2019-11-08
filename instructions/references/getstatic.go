package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/rtda/linker"
)

// Get static field from class
type GetStatic struct {
	base.Index16Instruction
	field *heap.Field
}

func (instr *GetStatic) Execute(frame *rtda.Frame) {
	if instr.field == nil {
		cp := frame.GetConstantPool()
		fieldRef := cp.GetConstantFieldRef(instr.Index)
		instr.field = linker.ResolveField(frame.GetBootLoader(), fieldRef, true)
	}

	class := instr.field.Class
	if class.InitializationNotStarted() {
		frame.RevertNextPC() // undo getstatic
		frame.Thread.InitClass(class)
		return
	}

	val := instr.field.GetStaticValue()
	frame.PushL(val, instr.field.IsLongOrDouble)
}
