package comparisons

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

func NewIfEQ() *IfCmp { return &IfCmp{cmpFn: eq} }
func NewIfNE() *IfCmp { return &IfCmp{cmpFn: ne} }
func NewIfLT() *IfCmp { return &IfCmp{cmpFn: lt} }
func NewIfLE() *IfCmp { return &IfCmp{cmpFn: le} }
func NewIfGT() *IfCmp { return &IfCmp{cmpFn: gt} }
func NewIfGE() *IfCmp { return &IfCmp{cmpFn: ge} }

func NewIfICmpEQ() *IfCmp { return &IfCmp{cmpFn: ieq} }
func NewIfICmpNE() *IfCmp { return &IfCmp{cmpFn: ine} }
func NewIfICmpLT() *IfCmp { return &IfCmp{cmpFn: ilt} }
func NewIfICmpLE() *IfCmp { return &IfCmp{cmpFn: ile} }
func NewIfICmpGT() *IfCmp { return &IfCmp{cmpFn: igt} }
func NewIfICmpGE() *IfCmp { return &IfCmp{cmpFn: ige} }
func NewIfACmpEQ() *IfCmp { return &IfCmp{cmpFn: aeq} }
func NewIfACmpNE() *IfCmp { return &IfCmp{cmpFn: ane} }

// extended, put here for convenience
func NewIfNull() *IfCmp    { return &IfCmp{cmpFn: null} }
func NewIfNonNull() *IfCmp { return &IfCmp{cmpFn: nonNull} }

func eq(frame *rtda.Frame) bool { return frame.PopInt() == 0 }
func ne(frame *rtda.Frame) bool { return frame.PopInt() != 0 }
func lt(frame *rtda.Frame) bool { return frame.PopInt() < 0 }
func le(frame *rtda.Frame) bool { return frame.PopInt() <= 0 }
func gt(frame *rtda.Frame) bool { return frame.PopInt() > 0 }
func ge(frame *rtda.Frame) bool { return frame.PopInt() >= 0 }

func ieq(frame *rtda.Frame) bool { return frame.PopInt() == frame.PopInt() }
func ine(frame *rtda.Frame) bool { return frame.PopInt() != frame.PopInt() }
func ilt(frame *rtda.Frame) bool { return frame.PopInt() > frame.PopInt() }
func ile(frame *rtda.Frame) bool { return frame.PopInt() >= frame.PopInt() }
func igt(frame *rtda.Frame) bool { return frame.PopInt() < frame.PopInt() }
func ige(frame *rtda.Frame) bool { return frame.PopInt() <= frame.PopInt() }
func aeq(frame *rtda.Frame) bool { return frame.PopRef() == frame.PopRef() }
func ane(frame *rtda.Frame) bool { return frame.PopRef() != frame.PopRef() }

func null(frame *rtda.Frame) bool    { return frame.PopRef() == nil }
func nonNull(frame *rtda.Frame) bool { return frame.PopRef() != nil }

// Branch if int or reference comparison succeeds
type IfCmp struct {
	base.BranchInstruction
	cmpFn func(frame *rtda.Frame) bool
}

func (instr *IfCmp) Execute(frame *rtda.Frame) {
	if instr.cmpFn(frame) {
		base.Branch(frame, instr.Offset)
	}
}
