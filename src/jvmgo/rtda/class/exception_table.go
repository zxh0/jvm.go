package class

import (
    //"fmt"
    cf "jvmgo/classfile"
)

type ExceptionTable struct {
    handlers []*ExceptionHandler
}

type ExceptionHandler struct {
    startPc     uint16
    endPc       uint16
    handlerPc   uint16
    catchType   *ConstantClass
}

func (self *ExceptionTable) copyExceptionTable(entries []*cf.ExceptionTableEntry, rtCp *ConstantPool) {
    self.handlers = make([]*ExceptionHandler, len(entries))
    for i, entry := range entries {
        self.handlers[i] = newExceptionHandler(entry, rtCp)
    }
}

func newExceptionHandler(entry *cf.ExceptionTableEntry, rtCp *ConstantPool) (*ExceptionHandler) {
    handler := &ExceptionHandler{}
    handler.startPc = entry.StartPc()
    handler.endPc = entry.EndPc()
    handler.handlerPc = entry.HandlerPc()
    catchType := uint(entry.CatchType())
    if catchType == 0 {
        handler.catchType = nil // catch all
    } else {
        handler.catchType = rtCp.GetConstant(catchType).(*ConstantClass)
    }
    return handler
}

func (self *ExceptionTable) FindExceptionHandler(exClass *Class) (*ExceptionHandler) {
    for _, handler := range self.handlers {
        if handler.catchType == nil || handler.catchType.Class() == exClass {
            return handler
        }
    }
    return nil
}
