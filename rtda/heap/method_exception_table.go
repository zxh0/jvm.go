package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type ExceptionHandler struct {
	startPc   int
	endPc     int
	handlerPc int
	catchType *ConstantClass
}

func (et *ExceptionHandler) HandlerPc() int {
	return et.handlerPc
}

type ExceptionTable struct {
	handlers []*ExceptionHandler
}

func (et *ExceptionTable) copyExceptionTable(entries []classfile.ExceptionTableEntry, rtCp *ConstantPool) {
	et.handlers = make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		et.handlers[i] = newExceptionHandler(entry, rtCp)
	}
}

func newExceptionHandler(entry classfile.ExceptionTableEntry, rtCp *ConstantPool) *ExceptionHandler {
	handler := &ExceptionHandler{}
	handler.startPc = int(entry.StartPc)
	handler.endPc = int(entry.EndPc)
	handler.handlerPc = int(entry.HandlerPc)
	catchType := uint(entry.CatchType)
	if catchType == 0 {
		handler.catchType = nil // catch all
	} else {
		handler.catchType = rtCp.GetConstant(catchType).(*ConstantClass)
	}
	return handler
}

func (et *ExceptionTable) FindExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range et.handlers {
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
