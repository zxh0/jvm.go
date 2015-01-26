package instructions

import (
    //"fmt"
    "jvmgo/rtda"
)

// Access jump table by index and jump
type tableswitch struct {
    // todo
}

func (self *tableswitch) fetchOperands(bcr *BytecodeReader) {
    // skip padding
    for bcr.pc % 4 != 0 {
        bcr.readUint8()
    }

    

    // todo
    panic("tableswitch")
}

func (self *tableswitch) Execute(thread *rtda.Thread) {
    // todo
    panic("tableswitch")
}
