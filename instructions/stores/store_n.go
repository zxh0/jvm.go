package stores

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

func NewStoreN(n uint, d bool) *StoreN {
	return &StoreN{n: n, d: d}
}

// xstore_n: Store XXX into local variable
type StoreN struct {
	base.NoOperandsInstruction
	n uint
	d bool // long or double
}

func (instr *StoreN) Execute(frame *rtda.Frame) {
	frame.Store(instr.n, instr.d)
}
