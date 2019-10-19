package conversions

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

func NewI2B() *X2Y { return &X2Y{castFn: i2b} }
func NewI2C() *X2Y { return &X2Y{castFn: i2c} }
func NewI2S() *X2Y { return &X2Y{castFn: i2s} }
func NewI2L() *X2Y { return &X2Y{castFn: i2l} }
func NewI2F() *X2Y { return &X2Y{castFn: i2f} }
func NewI2D() *X2Y { return &X2Y{castFn: i2d} }
func NewL2I() *X2Y { return &X2Y{castFn: l2i} }
func NewL2F() *X2Y { return &X2Y{castFn: l2f} }
func NewL2D() *X2Y { return &X2Y{castFn: l2d} }
func NewF2I() *X2Y { return &X2Y{castFn: f2i} }
func NewF2L() *X2Y { return &X2Y{castFn: f2l} }
func NewF2D() *X2Y { return &X2Y{castFn: f2d} }
func NewD2I() *X2Y { return &X2Y{castFn: d2i} }
func NewD2L() *X2Y { return &X2Y{castFn: d2l} }
func NewD2F() *X2Y { return &X2Y{castFn: d2f} }

func i2b(frame *rtda.Frame) { frame.PushInt(int32(int8(frame.PopInt()))) }
func i2c(frame *rtda.Frame) { frame.PushInt(int32(uint16(frame.PopInt()))) }
func i2s(frame *rtda.Frame) { frame.PushInt(int32(int16(frame.PopInt()))) }
func i2l(frame *rtda.Frame) { frame.PushLong(int64(frame.PopInt())) }
func i2f(frame *rtda.Frame) { frame.PushFloat(float32(frame.PopInt())) }
func i2d(frame *rtda.Frame) { frame.PushDouble(float64(frame.PopInt())) }
func l2i(frame *rtda.Frame) { frame.PushInt(int32(frame.PopLong())) }
func l2f(frame *rtda.Frame) { frame.PushFloat(float32(frame.PopLong())) }
func l2d(frame *rtda.Frame) { frame.PushDouble(float64(frame.PopLong())) }
func f2i(frame *rtda.Frame) { frame.PushInt(int32(frame.PopFloat())) }
func f2l(frame *rtda.Frame) { frame.PushLong(int64(frame.PopFloat())) }
func f2d(frame *rtda.Frame) { frame.PushDouble(float64(frame.PopFloat())) }
func d2i(frame *rtda.Frame) { frame.PushInt(int32(frame.PopDouble())) }
func d2l(frame *rtda.Frame) { frame.PushLong(int64(frame.PopDouble())) }
func d2f(frame *rtda.Frame) { frame.PushFloat(float32(frame.PopDouble())) }

type X2Y struct {
	base.NoOperandsInstruction
	castFn func(frame *rtda.Frame)
}

func (instr *X2Y) Execute(frame *rtda.Frame) {
	instr.castFn(frame)
}
