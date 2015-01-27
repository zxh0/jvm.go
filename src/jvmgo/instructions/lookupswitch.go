package instructions

import "jvmgo/rtda"

/*
lookupswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
npairs1
npairs2
npairs3
npairs4
match-offset pairs...
*/
// Access jump table by key match and jump
type lookupswitch struct {
    defaultOffset   int32
    npairs          int32
    // todo
}

func (self *lookupswitch) fetchOperands(bcr *BytecodeReader) {
    // skip padding
    for bcr.pc % 4 != 0 {
        bcr.readUint8()
    }

    // todo
    panic("tableswitch")
}

func (self *lookupswitch) Execute(thread *rtda.Thread) {
    // todo
    panic("lookupswitch")
}
