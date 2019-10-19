package math

import (
	"math"

	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

func NewDNeg() *DNeg { return &DNeg{} }
func NewDAdd() *DOp  { return &DOp{op: dadd} }
func NewDSub() *DOp  { return &DOp{op: dsub} }
func NewDMul() *DOp  { return &DOp{op: dmul} }
func NewDDiv() *DOp  { return &DOp{op: ddiv} }
func NewDRem() *DOp  { return &DOp{op: drem} }

func dadd(a, b float64) float64 { return a + b }
func dsub(a, b float64) float64 { return a - b }
func dmul(a, b float64) float64 { return a * b }
func ddiv(a, b float64) float64 { return a / b }
func drem(a, b float64) float64 { return math.Mod(a, b) } // todo

type DOp struct {
	base.NoOperandsInstruction
	op func(a, b float64) float64
}

func (instr *DOp) Execute(frame *rtda.Frame) {
	v2 := frame.PopDouble()
	v1 := frame.PopDouble()
	frame.PushDouble(instr.op(v1, v2))
}

// Negate double
type DNeg struct{ base.NoOperandsInstruction }

func (instr *DNeg) Execute(frame *rtda.Frame) {
	val := frame.PopDouble()
	frame.PushDouble(-val)
}
