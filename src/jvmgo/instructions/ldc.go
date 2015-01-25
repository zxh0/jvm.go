package instructions

import (
    "fmt"
    "unicode/utf8"
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
    case string: // todo
        // new string
        stringClass := frame.Method().Class().ClassLoader().StringClass()
        jStr := stringClass.NewObj()
        stack.PushRef(jStr)
        // init string
        codePoints := string2CodePoints(c.(string))
        //public String(int[] codePoints, int offset, int count)
        initMethod := stringClass.GetMethod("<init>", "([III)V")
        newFrame := rtda.NewFrame(initMethod)
        localVars := newFrame.LocalVars()
        localVars.SetRef(0, jStr) // this
        localVars.SetRef(1, rtc.NewIntArray(codePoints))
        localVars.SetInt(2, 0)
        localVars.SetInt(3, int32(len(codePoints)))
        thread.PushFrame(newFrame)
    case *rtc.ConstantClass: // todo
        class := c.(*rtc.ConstantClass).Class()
        stack.PushRef(class.Obj())
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

func string2CodePoints(str string) ([]rune) {
    runeCount := utf8.RuneCountInString(str)
    codePoints := make([]rune, runeCount)
    i := 0
    for len(str) > 0 {
        r, size := utf8.DecodeRuneInString(str)
        codePoints[i] = r
        i++
        str = str[size:]
    }
    return codePoints
}
