package stores

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

func NewStore(d bool) *Store {
	return &Store{d: d}
}

// xstore: Store XXX into local variable
type Store struct {
	base.Index8Instruction
	d bool // long or double
}

func (instr *Store) Execute(frame *rtda.Frame) {
	frame.Store(instr.Index, instr.d)
}
