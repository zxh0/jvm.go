package constants

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func NewConstNull() *Const            { return &Const{k: heap.EmptySlot, d: false} }
func NewConstInt(n int32) *Const      { return &Const{k: heap.NewIntSlot(n), d: false} }
func NewConstLong(n int64) *Const     { return &Const{k: heap.NewLongSlot(n), d: true} }
func NewConstFloat(n float32) *Const  { return &Const{k: heap.NewFloatSlot(n), d: false} }
func NewConstDouble(n float64) *Const { return &Const{k: heap.NewDoubleSlot(n), d: true} }

// xconst: Push XXX
type Const struct {
	base.NoOperandsInstruction
	k heap.Slot
	d bool // long or double
}

func (instr *Const) Execute(frame *rtda.Frame) {
	frame.PushL(instr.k, instr.d)
}
