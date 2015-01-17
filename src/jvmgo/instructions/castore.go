package instructions

import "jvmgo/rtda"

// Store into char array 
type castore struct {

}

func (self *castore) fetchOperands(bcr *BytecodeReader) {
    // no operands
}

func (self *castore) execute(thread *rtda.Thread) {
    // todo
}
