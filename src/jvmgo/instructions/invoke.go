package instructions

import (
    "jvmgo/rtda"
    "jvmgo/rtda/class"
)

// Invoke a class (static) method 
type invokestatic struct {Index16Instruction}
func (self *invokestatic) Execute(thread *rtda.Thread) {
    frame := thread.CurrentFrame()
    stack := frame.OperandStack()

    cp := frame.Method().Class().ConstantPool()
    cMethodRef := cp.GetConstant(self.index).(class.ConstantMethodref)
    field := cFieldRef.Field()
    // todo
}

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations 
type invokespecial struct {Index16Instruction}
func (self *invokespecial) Execute(thread *rtda.Thread) {
    // todo
}

// Invoke instance method; dispatch based on class
type invokevirtual struct {Index16Instruction}
func (self *invokevirtual) Execute(thread *rtda.Thread) {
    // todo
}

// Invoke interface method
type invokeinterface struct {
    index   uint16
    count   uint8
    // 0
}
func (self *invokeinterface) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
    self.count = bcr.readUint8()
    bcr.readUint8() // must be 0
}
func (self *invokeinterface) Execute(thread *rtda.Thread) {
    // todo
}

// Invoke dynamic method
type invokedynamic struct {
    index uint16
    // 0
    // 0
}
func (self *invokedynamic) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint16()
    bcr.readUint8() // must be 0
    bcr.readUint8() // must be 0
}
func (self *invokedynamic) Execute(thread *rtda.Thread) {
    // todo
}
