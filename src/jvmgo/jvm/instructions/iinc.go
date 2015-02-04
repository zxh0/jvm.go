package instructions

import "jvmgo/rtda"

// Increment local variable by constant
type iinc struct {
    index   uint8
    _const  int8
}

func (self *iinc) fetchOperands(bcr *BytecodeReader) {
    self.index = bcr.readUint8()
    self._const = bcr.readInt8()
}

func (self *iinc) Execute(frame *rtda.Frame) {
    localVars := frame.LocalVars()
    index := uint(self.index)
    val := localVars.GetInt(index)
    val += int32(self._const)
    localVars.SetInt(index, val)
}
