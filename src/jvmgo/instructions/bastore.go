package instructions

import "jvmgo/rtda"

// Store into byte or boolean array 
type bastore struct {

}

func (self *bastore) fetchOperands(bcr *BytecodeReader) {
    // no operands
}

func (self *bastore) execute(thread *rtda.Thread) {
    // todo
}
