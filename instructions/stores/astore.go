package stores

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func NewIAStore() *AStore { return &AStore{atype: heap.ATInt} }
func NewLAStore() *AStore { return &AStore{atype: heap.ATLong} }
func NewFAStore() *AStore { return &AStore{atype: heap.ATFloat} }
func NewDAStore() *AStore { return &AStore{atype: heap.ATDouble} }
func NewAAStore() *AStore { return &AStore{atype: 0} }
func NewBAStore() *AStore { return &AStore{atype: heap.ATByte} }
func NewCAStore() *AStore { return &AStore{atype: heap.ATChar} }
func NewSAStore() *AStore { return &AStore{atype: heap.ATShort} }

// xastore: Store into XXX array
type AStore struct {
	base.NoOperandsInstruction
	atype byte
}

func (instr *AStore) Execute(frame *rtda.Frame) {
	d := instr.atype == heap.ATLong || instr.atype == heap.ATDouble
	slot := frame.PopL(d)
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
		arrRef.GetBytes()[index] = int8(slot.IntValue())
	case heap.ATChar:
		arrRef.GetChars()[index] = uint16(slot.IntValue())
	case heap.ATShort:
		arrRef.GetShorts()[index] = int16(slot.IntValue())
	case heap.ATInt:
		arrRef.GetInts()[index] = slot.IntValue()
	case heap.ATLong:
		arrRef.GetLongs()[index] = slot.LongValue()
	case heap.ATFloat:
		arrRef.GetFloats()[index] = slot.FloatValue()
	case heap.ATDouble:
		arrRef.GetDoubles()[index] = slot.DoubleValue()
	default:
		arrRef.GetRefs()[index] = slot.Ref
	}
}
