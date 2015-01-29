package instructions

import (
    "jvmgo/rtda"
    rtclass "jvmgo/rtda/class"
)

// Create new object
type new_ struct {Index16Instruction}
func (self *new_) Execute(frame *rtda.Frame) {
    cp := frame.Method().Class().ConstantPool()
    cClass := cp.GetConstant(self.index).(*rtclass.ConstantClass)
    class := cClass.Class()

    if class.NotInitialized() {
        thread := frame.Thread()
        frame.SetNextPC(thread.PC())
        initClass(class, thread)
    } else {
        ref := class.NewObj()
        frame.OperandStack().PushRef(ref)
    }
}

// Create new array
type newarray struct {
    atype uint8
}
func (self *newarray) fetchOperands(bcr *BytecodeReader) {
    self.atype = bcr.readUint8()
}
func (self *newarray) Execute(frame *rtda.Frame) {
    stack := frame.OperandStack()
    count := stack.PopInt()
    ref := rtclass.NewPrimitiveArray(self.atype, count)
    stack.PushRef(ref)
}

// Create new array of reference
type anewarray struct {Index16Instruction}
func (self *anewarray) Execute(frame *rtda.Frame) {
    cp := frame.Method().Class().ConstantPool()
    cClass := cp.GetConstant(self.index).(*rtclass.ConstantClass)
    class := cClass.Class()
    
    // todo
    if class.NotInitialized() {
        //frame.SetNextPC(thread.PC())
        //initClass(class, thread)
        //panic("class not initialized!"  + class.Name())
    }

    stack := frame.OperandStack()
    count := stack.PopInt()
    ref := rtclass.NewRefArray(count)
    stack.PushRef(ref)
}

// Create new multidimensional array
type multianewarray struct {
    index       uint16
    dimensions  uint8
}
func (self *multianewarray) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
    self.dimensions = bcr.readUint8()
}
func (self *multianewarray) Execute(frame *rtda.Frame) {
    // todo
    panic("todo multianewarray")
}
