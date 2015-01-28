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

func (self *ExceptionTable) copyExceptionTable(entries []*cf.ExceptionTableEntry) {
    self.handlers = make([]*ExceptionHandler, len(entries))
    for i, entry := range entries {
        handler := &ExceptionHandler{}
        handler.startPc = entry.StartPc()
        handler.endPc = entry.EndPc()
        handler.handlerPc = entry.HandlerPc()
        handler.catchType = nil
        self.handlers[i] = handler
    }
}
