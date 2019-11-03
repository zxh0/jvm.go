package zip

import (
	"hash/crc32"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_crc(updateBytes, "updateBytes", "(I[BII)I")
}

func _crc(method native.Method, name, desc string) {
	native.Register("java/util/zip/CRC32", name, desc, method)
}

// private native static int updateBytes(int crc, byte[] b, int off, int len);
// (I[BII)I
func updateBytes(frame *rtda.Frame) {
	crc := uint32(frame.GetIntVar(0))
	byteArr := frame.GetRefVar(1)
	off := frame.GetIntVar(2)
	_len := frame.GetIntVar(3)

	goBytes := byteArr.GetGoBytes()
	goBytes = goBytes[off : off+_len]
	// func Update(crc uint32, tab *Table, p []byte) uint32
	crc = crc32.Update(crc, crc32.IEEETable, goBytes)

	frame.PushInt(int32(crc))
}
