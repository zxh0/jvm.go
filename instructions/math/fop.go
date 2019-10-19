package math

import (
	"math"

	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

func NewFNeg() *FNeg { return &FNeg{} }
func NewFAdd() *FOp  { return &FOp{op: fadd} }
func NewFSub() *FOp  { return &FOp{op: fsub} }
func NewFMul() *FOp  { return &FOp{op: fmul} }
func NewFDiv() *FOp  { return &FOp{op: fdiv} }
func NewFRem() *FOp  { return &FOp{op: frem} }

func fadd(a, b float32) float32 { return a + b }
func fsub(a, b float32) float32 { return a - b }
func fmul(a, b float32) float32 { return a * b }
func fdiv(a, b float32) float32 { return a / b }
func frem(a, b float32) float32 { return float32(math.Mod(float64(a), float64(b))) } // todo

type FOp struct {
	base.NoOperandsInstruction
	op func(a, b float32) float32
}

func (instr *FOp) Execute(frame *rtda.Frame) {
	v2 := frame.PopFloat()
	v1 := frame.PopFloat()
	frame.PushFloat(instr.op(v1, v2))
}

// Negate float
type FNeg struct{ base.NoOperandsInstruction }

func (instr *FNeg) Execute(frame *rtda.Frame) {
	val := frame.PopFloat()
	frame.PushFloat(-val)
}
