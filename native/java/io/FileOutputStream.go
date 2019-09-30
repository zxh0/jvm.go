package io

import (
	"os"

	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_fos(writeBytes, "writeBytes", "([BIIZ)V")
}

func _fos(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/io/FileOutputStream", name, desc, method)
}

// private native void writeBytes(byte b[], int off, int len, boolean append) throws IOException;
// ([BIIZ)V
func writeBytes(frame *rtda.Frame) {
	fosObj := frame.GetRefVar(0)     // this
	byteArrObj := frame.GetRefVar(1) // b
	offset := frame.GetIntVar(2)     // off
	length := frame.GetIntVar(3)     // len
	//frame.GetBooleanVar(4) // append

	fdObj := fosObj.GetFieldValue("fd", "Ljava/io/FileDescriptor;").Ref
	if fdObj.Extra == nil {
		goFd := fdObj.GetFieldValue("fd", "I").IntValue()
		switch goFd {
		case 0:
			fdObj.Extra = os.Stdin
		case 1:
			fdObj.Extra = os.Stdout
		case 2:
			fdObj.Extra = os.Stderr
		default:
			fdObj.Extra = os.Stdout
		}
	}
	goFile := fdObj.Extra.(*os.File)

	goBytes := byteArrObj.GoBytes()
	goBytes = goBytes[offset : offset+length]
	goFile.Write(goBytes)
}
