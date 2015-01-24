package instructions

import (
    "fmt"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// Push item from run-time constant pool 
type ldc struct {Index8Instruction}
func (self *ldc) Execute(thread *rtda.Thread) {
    _ldc(thread, self.index)
}

// Push item from run-time constant pool (wide index)
type ldc_w struct {Index16Instruction}
func (self *ldc_w) Execute(thread *rtda.Thread) {
    _ldc(thread, self.index)
}

func _ldc(thread *rtda.Thread, index uint) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()
    cp := frame.Method().Class().ConstantPool()
    c := cp.GetConstant(index)

    switch c.(type) {
    case int32: stack.PushInt(c.(int32))
    case float32: stack.PushFloat(c.(float32))
    case string: panic("todo ldc string!!")
    case *rtc.ConstantClass: 
        panic("todo ldc class")
        //class := c.(*rtc.ConstantClass).Class()
        //stack.PushRef(class.(*rtc.Obj))
    default: 
        fmt.Printf("CCC:::%v\n", c)
        panic("todo ldc!!!")
        // todo
        // ref to String
        // ref to Class
        // ref to MethodType or MethodHandle
    }
}

// Push long or double from run-time constant pool (wide index) 
type ldc2_w struct {Index16Instruction}
func (self *ldc2_w) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()
    cp := frame.Method().Class().ConstantPool()
    c := cp.GetConstant(self.index)

    switch c.(type) {
    case int64: stack.PushLong(c.(int64))
    case float64: stack.PushDouble(c.(float64))
    // todo
    }
}
