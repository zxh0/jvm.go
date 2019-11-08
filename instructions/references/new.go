package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/rtda/linker"
)

// Create new object
type New struct {
	base.Index16Instruction
	class *heap.Class
}

func (instr *New) Execute(frame *rtda.Frame) {
	if instr.class == nil {
		cp := frame.GetConstantPool()
		classRef := cp.GetConstantClass(instr.Index)
		instr.class = linker.ResolveClass(frame.GetBootLoader(), classRef)
	}

	// init class
	if instr.class.InitializationNotStarted() {
		frame.RevertNextPC() // undo new
		frame.Thread.InitClass(instr.class)
		return
	}

	ref := instr.class.NewObj()
	frame.PushRef(ref)
}
