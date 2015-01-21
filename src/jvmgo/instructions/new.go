package instructions

import (
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

// Create new object
type new_ struct {Index16Instruction}
func (self *new_) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()
    cp := frame.Method().Class().ConstantPool()
    cClass := cp.GetConstant(self.index).(class.ConstantClass)
    class := cClass.Class()
    ref := class.NewObj()
    stack.PushRef(ref)
}

// Create new array
type newarray struct {
    atype uint8
}
func (self *newarray) fetchOperands(bcr *BytecodeReader) {
    self.atype = bcr.readUint8()
}
func (self *newarray) Execute(thread *rtda.Thread) {
    stack := thread.CurrentFrame().OperandStack()
    count := stack.PopInt()
    ref := class.NewArray(self.atype, count)
    stack.PushRef(ref)
}

// Create new array of reference
type anewarray struct {Index16Instruction}
func (self *anewarray) Execute(thread *rtda.Thread) {
    // todo
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
func (self *multianewarray) Execute(thread *rtda.Thread) {
    // todo
}
