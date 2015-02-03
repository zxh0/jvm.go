package instructions

import (
    "fmt"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// Push item from run-time constant pool 
type ldc struct {Index8Instruction}
func (self *ldc) Execute(frame *rtda.Frame) {
    _ldc(frame, self.index)
}

// Push item from run-time constant pool (wide index)
type ldc_w struct {Index16Instruction}
func (self *ldc_w) Execute(frame *rtda.Frame) {
    _ldc(frame, self.index)
}

func _ldc(frame *rtda.Frame, index uint) {
    stack := frame.OperandStack()
    cp := frame.Method().Class().ConstantPool()
    c := cp.GetConstant(index)

    switch c.(type) {
    case int32: stack.PushInt(c.(int32))
    case float32: stack.PushFloat(c.(float32))
    case *rtc.ConstantString: 
        constStr := c.(*rtc.ConstantString)
        if constStr.JStr() == nil {
            chars, jStr := rtda.NewJString(constStr.GoStr(), frame)
            jStr = rtc.InternString(chars, jStr)
            constStr.SetJStr(jStr)
        }
        stack.PushRef(constStr.JStr())
    case *rtc.ConstantClass: // todo
        class := c.(*rtc.ConstantClass).Class()
        stack.PushRef(class.JClass())
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
func (self *ldc2_w) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    cp := frame.Method().Class().ConstantPool()
    c := cp.GetConstant(self.index)

    switch c.(type) {
    case int64: stack.PushLong(c.(int64))
    case float64: stack.PushDouble(c.(float64))
    // todo
    }
}
