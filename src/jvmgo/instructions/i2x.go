package instructions

import "jvmgo/rtda"

// Convert int to byte
type i2b struct {NoOperandsInstruction}
func (self *i2b) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    i := stack.PopInt()
    b := int32(int8(i))
    stack.PushInt(b)
}

// Convert int to char
type i2c struct {NoOperandsInstruction}
func (self *i2c) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    i := stack.PopInt()
    c := int32(uint16(i))
    stack.PushInt(c)
}

// Convert int to short
type i2s struct {NoOperandsInstruction}
func (self *i2s) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    i := stack.PopInt()
    s := int32(int16(i))
    stack.PushInt(s)
}

// Convert int to long
type i2l struct {NoOperandsInstruction}
func (self *i2l) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    i := stack.PopInt()
    l := int64(i)
    stack.PushLong(l)
}

// Convert int to float
type i2f struct {NoOperandsInstruction}
func (self *i2f) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    i := stack.PopInt()
    f := float32(i)
    stack.PushFloat(f)
}

// Convert int to double
type i2d struct {NoOperandsInstruction}
func (self *i2d) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    i := stack.PopInt()
    d := float64(i)
    stack.PushDouble(d)
}
