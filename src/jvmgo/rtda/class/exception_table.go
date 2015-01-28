package class

import cf "jvmgo/classfile"

type ExceptionTable struct {
    handlers []*ExceptionHandler
}

type ExceptionHandler struct {
    startPc     uint16
    endPc       uint16
    handlerPc   uint16
    catchType   *ConstantClass
}

func (self *ExceptionTable) copy(entries []*cf.ExceptionTableEntry) {
    self.handlers = make([]*ExceptionHandler, len(entries))
    // for i, entry := range entries {
    //     handler := &ExceptionHandler{}
    //     handler.startPc = entry.startPc
    //     handler.endPc = entry.endPc
    //     handler.handlerPc = entry.handlerPc
    //     handler.catchType = nil
    //     self.handlers[i] = handler
    // }
}
