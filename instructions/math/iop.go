package math

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

func NewIInc() *IInc { return &IInc{} }
func NewINeg() *INeg { return &INeg{} }
func NewIAdd() *IOp  { return &IOp{op: iadd} }
func NewISub() *IOp  { return &IOp{op: isub} }
func NewIMul() *IOp  { return &IOp{op: imul} }
func NewIDiv() *IOp  { return &IOp{op: idiv, div: true} }
func NewIRem() *IOp  { return &IOp{op: irem, div: true} }
func NewIShl() *IOp  { return &IOp{op: ishl} }
func NewIShr() *IOp  { return &IOp{op: ishr} }
func NewIUShr() *IOp { return &IOp{op: iushr} }
func NewIAnd() *IOp  { return &IOp{op: iand} }
func NewIOr() *IOp   { return &IOp{op: ior} }
func NewIXor() *IOp  { return &IOp{op: ixor} }

func iadd(a, b int32) int32  { return a + b }
func isub(a, b int32) int32  { return a - b }
func imul(a, b int32) int32  { return a * b }
func idiv(a, b int32) int32  { return a / b }
func irem(a, b int32) int32  { return a % b }
func iand(a, b int32) int32  { return a & b }
func ior(a, b int32) int32   { return a | b }
func ixor(a, b int32) int32  { return a ^ b }
func ishl(a, b int32) int32  { return a << (b & 0x1f) }
func ishr(a, b int32) int32  { return a >> (b & 0x1f) }
func iushr(a, b int32) int32 { return int32(uint32(a) >> (b & 0x1f)) }

type IOp struct {
	base.NoOperandsInstruction
	op  func(a, b int32) int32
	div bool
}

func (instr *IOp) Execute(frame *rtda.Frame) {
	v2 := frame.PopInt()
	v1 := frame.PopInt()
	if instr.div && v2 == 0 {
		frame.Thread.ThrowDivByZero()
	} else {
		frame.PushInt(instr.op(v1, v2))
	}
}

// Negate int
type INeg struct{ base.NoOperandsInstruction }

func (instr *INeg) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	frame.PushInt(-val)
}

// Increment local variable by constant
type IInc struct {
	Index uint
	Const int32
}

func (instr *IInc) FetchOperands(reader *base.CodeReader) {
	instr.Index = uint(reader.ReadUint8())
	instr.Const = int32(reader.ReadInt8())
}

func (instr *IInc) Execute(frame *rtda.Frame) {
	val := frame.GetIntVar(instr.Index)
	val += instr.Const
	frame.SetIntVar(instr.Index, val)
}
