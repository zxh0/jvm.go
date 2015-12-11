package heap

import (
	cf "github.com/zxh0/jvm.go/jvmgo/classfile"
)

type ExceptionHandler struct {
	startPc   int
	endPc     int
	handlerPc int
	catchType *ConstantClass
}

func (self *ExceptionHandler) HandlerPc() int {
	return self.handlerPc
}

type ExceptionTable struct {
	handlers []*ExceptionHandler
}

func (self *ExceptionTable) copyExceptionTable(entries []*cf.ExceptionTableEntry, rtCp *ConstantPool) {
	self.handlers = make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		self.handlers[i] = newExceptionHandler(entry, rtCp)
	}
}

func newExceptionHandler(entry *cf.ExceptionTableEntry, rtCp *ConstantPool) *ExceptionHandler {
	handler := &ExceptionHandler{}
	handler.startPc = int(entry.StartPc())
	handler.endPc = int(entry.EndPc())
	handler.handlerPc = int(entry.HandlerPc())
	catchType := uint(entry.CatchType())
	if catchType == 0 {
		handler.catchType = nil // catch all
	} else {
		handler.catchType = rtCp.GetConstant(catchType).(*ConstantClass)
	}
	return handler
}

func (self *ExceptionTable) FindExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range self.handlers {
		// jvms: The start_pc is inclusive and end_pc is exclusive
		if pc >= handler.startPc && pc < handler.endPc {
			if handler.catchType == nil || // catch all
				handler.catchType.Class() == exClass ||
				handler.catchType.Class().isSuperClassOf(exClass) {

				return handler
			}
		}
	}
	return nil
}
