package loads

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func NewIALoad() *ALoad { return &ALoad{atype: heap.ATInt} }
func NewLALoad() *ALoad { return &ALoad{atype: heap.ATLong} }
func NewFALoad() *ALoad { return &ALoad{atype: heap.ATFloat} }
func NewDALoad() *ALoad { return &ALoad{atype: heap.ATDouble} }
func NewAALoad() *ALoad { return &ALoad{atype: 0} }
func NewBALoad() *ALoad { return &ALoad{atype: heap.ATByte} }
func NewCALoad() *ALoad { return &ALoad{atype: heap.ATChar} }
func NewSALoad() *ALoad { return &ALoad{atype: heap.ATShort} }

// xaload: Load XXX from array
type ALoad struct {
	base.NoOperandsInstruction
	atype byte
}

func (instr *ALoad) Execute(frame *rtda.Frame) {
	index := frame.PopInt()
	arrRef := frame.PopRef()

	if arrRef == nil {
		frame.Thread.ThrowNPE()
		return
	}
	if index < 0 || index >= arrRef.ArrayLength() {
		frame.Thread.ThrowArrayIndexOutOfBoundsException(index)
		return
	}

	switch instr.atype {
	case heap.ATByte:
		frame.PushInt(int32(arrRef.GetBytes()[index]))
	case heap.ATChar:
		frame.PushInt(int32(arrRef.GetChars()[index]))
	case heap.ATShort:
		frame.PushInt(int32(arrRef.GetShorts()[index]))
	case heap.ATInt:
		frame.PushInt(arrRef.GetInts()[index])
	case heap.ATLong:
		frame.PushLong(arrRef.GetLongs()[index])
	case heap.ATFloat:
		frame.PushFloat(arrRef.GetFloats()[index])
	case heap.ATDouble:
		frame.PushDouble(arrRef.GetDoubles()[index])
	default:
		frame.PushRef(arrRef.GetRefs()[index])
	}
}
