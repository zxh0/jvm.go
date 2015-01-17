package instructions

import "jvmgo/rtda"

// Load reference from local variable 
type aload struct {
    index uint8
}

func (self *aload) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint8()
}

func (self *aload) execute(thread *rtda.Thread) {
    ref := thread.CurrentFrame().LocalVars().GetRef(uint(self.index))
    thread.CurrentFrame().OperandStack().PushRef(ref)
}
