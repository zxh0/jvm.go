package rtda

import (
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

var (
	_native_hack_ireturn = []byte{0xfe, 0xac}
	_native_hack_lreturn = []byte{0xfe, 0xad}
	_native_hack_freturn = []byte{0xfe, 0xae}
	_native_hack_dreturn = []byte{0xfe, 0xaf}
	_native_hack_areturn = []byte{0xfe, 0xb0}
	_native_hack_return  = []byte{0xfe, 0xb1}
)

func newNativeFrame(thread *Thread, method *heap.Method) *Frame {
	frame := &Frame{}
	frame.thread = thread
	frame.method = method
	frame.localVars = newLocalVars(method.ArgSlotCount()) // todo
	frame.operandStack = newOperandStack(4)               // todo

	code := method.Code()
	if code == nil {
		code = getHackCode(method.Descriptor())
		method.HackSetCode(code)
	}

	return frame
}

func getHackCode(methodDescriptor string) []byte {
	rd := heap.GetReturnDescriptor(methodDescriptor)
	switch rd[0] {
	case 'V':
		return _native_hack_return
	case 'L', '[':
		return _native_hack_areturn
	case 'D':
		return _native_hack_dreturn
	case 'F':
		return _native_hack_freturn
	case 'J':
		return _native_hack_lreturn
	default:
		return _native_hack_ireturn
	}
}
