package instructions

import "jvmgo/rtda"

// Invoke a class (static) method 
type invokestatic struct {Index16Instruction}
func (self *invokestatic) execute(thread *rtda.Thread) {
    // todo
}

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations 
type invokespecial struct {Index16Instruction}
func (self *invokespecial) execute(thread *rtda.Thread) {
    // todo
}

// Invoke instance method; dispatch based on class
type invokevirtual struct {Index16Instruction}
func (self *invokevirtual) execute(thread *rtda.Thread) {
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
func (self *invokeinterface) execute(thread *rtda.Thread) {
    // todo
}
