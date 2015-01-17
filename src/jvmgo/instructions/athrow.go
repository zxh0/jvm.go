package instructions

import "jvmgo/rtda"

// Throw exception or error
type athrow struct {
    
}

func (self *athrow) fetchOperands(bcr *BytecodeReader) {
    
}

func (self *athrow) execute(thread *rtda.Thread) {
    // todo
}
