package instructions

import "jvmgo/jvm/rtda"

// Jump subroutine
type jsr struct{ BranchInstruction }

func (self *jsr) Execute(frame *rtda.Frame) {
	panic("todo")
}
