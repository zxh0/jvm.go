package math

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

func NewLNeg() *LNeg   { return &LNeg{} }
func NewLAdd() *LOp    { return &LOp{op: ladd} }
func NewLSub() *LOp    { return &LOp{op: lsub} }
func NewLMul() *LOp    { return &LOp{op: lmul} }
func NewLDiv() *LOp    { return &LOp{op: ldiv, div: true} }
func NewLRem() *LOp    { return &LOp{op: lrem, div: true} }
func NewLShl() *LShOp  { return &LShOp{op: lshl} }
func NewLShr() *LShOp  { return &LShOp{op: lshr} }
func NewLUShr() *LShOp { return &LShOp{op: lushr} }
func NewLAnd() *LOp    { return &LOp{op: land} }
func NewLOr() *LOp     { return &LOp{op: lor} }
func NewLXor() *LOp    { return &LOp{op: lxor} }

func ladd(a, b int64) int64        { return a + b }
func lsub(a, b int64) int64        { return a - b }
func lmul(a, b int64) int64        { return a * b }
func ldiv(a, b int64) int64        { return a / b }
func lrem(a, b int64) int64        { return a % b }
func land(a, b int64) int64        { return a & b }
func lor(a, b int64) int64         { return a | b }
func lxor(a, b int64) int64        { return a ^ b }
func lshl(a int64, b int32) int64  { return a << (b & 0x3f) }
func lshr(a int64, b int32) int64  { return a >> (b & 0x3f) }
func lushr(a int64, b int32) int64 { return int64(uint64(a) >> (b & 0x3f)) }

type LOp struct {
	base.NoOperandsInstruction
	op  func(a, b int64) int64
	div bool
}

func (instr *LOp) Execute(frame *rtda.Frame) {
	v2 := frame.PopLong()
	v1 := frame.PopLong()
	if instr.div && v2 == 0 {
		frame.Thread.ThrowDivByZero()
	} else {
		frame.PushLong(instr.op(v1, v2))
	}
}

type LShOp struct {
	base.NoOperandsInstruction
	op func(a int64, b int32) int64
}

func (instr *LShOp) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopLong()
	frame.PushLong(instr.op(v1, v2))
}

// Negate long
type LNeg struct{ base.NoOperandsInstruction }

func (instr *LNeg) Execute(frame *rtda.Frame) {
	val := frame.PopLong()
	frame.PushLong(-val)
}
