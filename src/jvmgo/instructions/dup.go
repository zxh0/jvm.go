package instructions

import (
    . "jvmgo/any"
    "jvmgo/rtda"
)

// Duplicate the top operand stack value
type dup struct {}
func (self *dup) fetchOperands(bcr *BytecodeReader) {}
func (self *dup) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val := stack.Pop()
    stack.Push(val)
    stack.Push(val)
}

// Duplicate the top operand stack value and insert two values down
type dup_x1 struct {}
func (self *dup_x1) fetchOperands(bcr *BytecodeReader) {}
func (self *dup_x1) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val1 := stack.Pop()
    val2 := stack.Pop()
    stack.Push(val1)
    stack.Push(val2)
    stack.Push(val1)
}

// Duplicate the top operand stack value and insert two or three values down 
type dup_x2 struct {}
func (self *dup_x2) fetchOperands(bcr *BytecodeReader) {}
func (self *dup_x2) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val1 := stack.Pop()
    val2 := stack.Pop()
    if isLongOrDouble(val2) {
        // form2
        stack.Push(val1)
        stack.Push(val2)
        stack.Push(val1)
    } else {
        // form1
        val3 := stack.Pop()
        stack.Push(val1)
        stack.Push(val3)
        stack.Push(val2)
        stack.Push(val1)
    }
}

// Duplicate the top one or two operand stack values 
type dup2 struct {}
func (self *dup2) fetchOperands(bcr *BytecodeReader) {}
func (self *dup2) execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    val1 := stack.Pop()
    if isLongOrDouble(val1) {
        stack.Push(val1)
        stack.Push(val1)
    } else {
        // form1
        val2 := stack.Pop()
        stack.Push(val2)
        stack.Push(val1)
        stack.Push(val2)
        stack.Push(val1)
    }
}

func isLongOrDouble(x Any) (bool) {
    if _, ok := x.(int64); ok {
        return true
    }
    if _, ok := x.(float64); ok {
        return true
    }
    return false
}
