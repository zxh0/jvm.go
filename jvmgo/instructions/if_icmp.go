package instructions

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda"
)

// Branch if int comparison succeeds
type if_icmpeq struct{ BranchInstruction }

func (self *if_icmpeq) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		branch(frame, self.offset)
	}
}

type if_icmpne struct{ BranchInstruction }

func (self *if_icmpne) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		branch(frame, self.offset)
	}
}

type if_icmplt struct{ BranchInstruction }

func (self *if_icmplt) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		branch(frame, self.offset)
	}
}

type if_icmple struct{ BranchInstruction }

func (self *if_icmple) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		branch(frame, self.offset)
	}
}

type if_icmpgt struct{ BranchInstruction }

func (self *if_icmpgt) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		branch(frame, self.offset)
	}
}

type if_icmpge struct{ BranchInstruction }

func (self *if_icmpge) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		branch(frame, self.offset)
	}
}

func _icmpPop(frame *rtda.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}
