package instructions

import "jvmgo/rtda"

// Store into reference array 
type aastore struct {

}

func (self *aastore) fetchOperands(bcr *BytecodeReader) {
    // no operands
}

func (self *aastore) execute(thread *rtda.Thread) {
    // todo
}
