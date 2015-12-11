package io

import (
	"io"
	"net"
	"time"

	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_sis(sis_init, "init", "()V")
	_sis(sis_socketRead0, "socketRead0", "(Ljava/io/FileDescriptor;[BIII)I")
}

func _sis(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/net/SocketInputStream", name, desc, method)
}

// native java/net/SocketOutputStream~init~()V
func sis_init(frame *rtda.Frame) {
}

//private native int socketRead0(FileDescriptor fd, byte b[], int off, int len, int timeout)
// java/net/SocketInputStream~socketRead0~(Ljava/io/FileDescriptor;[BIII)I
func sis_socketRead0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	//this := vars.GetThis()
	fd := vars.GetRef(1)
	buf := vars.GetRef(2)
	off := vars.GetInt(3)
	_len := vars.GetInt(4)

	conn := fd.Extra().(net.Conn)

	_timeout := vars.GetInt(5)
	if _timeout > 0 {
		conn.SetDeadline(time.Now().Add(time.Duration(_timeout) * time.Millisecond))
	}

	goBuf := buf.GoBytes()
	goBuf = goBuf[off : off+_len]

	n, err := conn.Read(goBuf)
	if err == nil || n > 0 || err == io.EOF {
		frame.OperandStack().PushInt(int32(n))
	} else {
		// todo
		panic(err.Error())
		//frame.Thread().ThrowIOException(err.Error())
	}

}
