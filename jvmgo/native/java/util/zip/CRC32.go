package zip

import (
	"hash/crc32"

	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_crc(updateBytes, "updateBytes", "(I[BII)I")
}

func _crc(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/util/zip/CRC32", name, desc, method)
}

// private native static int updateBytes(int crc, byte[] b, int off, int len);
// (I[BII)I
func updateBytes(frame *rtda.Frame) {
	vars := frame.LocalVars()
	crc := uint32(vars.GetInt(0))
	byteArr := vars.GetRef(1)
	off := vars.GetInt(2)
	_len := vars.GetInt(3)

	goBytes := byteArr.GoBytes()
	goBytes = goBytes[off : off+_len]
	// func Update(crc uint32, tab *Table, p []byte) uint32
	crc = crc32.Update(crc, crc32.IEEETable, goBytes)

	stack := frame.OperandStack()
	stack.PushInt(int32(crc))
}
