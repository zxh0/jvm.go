package instructions

import (
    //"fmt"
    "jvmgo/rtda"
    "jvmgo/rtda/class" // rtc
)

// Fetch field from object
type getfield struct {Index16Instruction}
func (self *getfield) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()
    ref := stack.PopRef()
    if ref == nil {
        // todo NullPointerException
        panic("NPE")
    }

    cp := frame.Method().Class().ConstantPool()
    cFieldRef := cp.GetConstant(self.index).(*class.ConstantFieldref)
    field := cFieldRef.InstanceField()
    val := field.GetValue(ref)

    stack.Push(val)
}

// Get static field from class 
type getstatic struct {Index16Instruction}
func (self *getstatic) Execute(thread *rtda.Thread) {
    currentFrame := thread.CurrentFrame()
    currentMethod := currentFrame.Method()
    currentClass := currentMethod.Class()

    cp := currentClass.ConstantPool()
    cFieldRef := cp.GetConstant(self.index).(*class.ConstantFieldref)
    field := cFieldRef.StaticField()

    classOfField := field.Class()
    if classOfField.NotInitialized() {
        if classOfField != currentClass || !currentMethod.IsClinit() {
            currentFrame.SetNextPC(thread.PC()) // undo getstatic
            initClass(field.Class(), thread)
            return
        }
    }

    val := field.GetStaticValue()
    currentFrame.OperandStack().Push(val)
}
