package io

import (
	"io"
	"net"
	"time"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_sis(sis_init, "init", "()V")
	_sis(sis_socketRead0, "socketRead0", "(Ljava/io/FileDescriptor;[BIII)I")
}

func _sis(method native.Method, name, desc string) {
	native.Register("java/net/SocketInputStream", name, desc, method)
}

// native java/net/SocketOutputStream~init~()V
func sis_init(frame *rtda.Frame) {
}

//private native int socketRead0(FileDescriptor fd, byte b[], int off, int len, int timeout)
// java/net/SocketInputStream~socketRead0~(Ljava/io/FileDescriptor;[BIII)I
func sis_socketRead0(frame *rtda.Frame) {
	//this := frame.GetThis()
	fd := frame.GetRefVar(1)
	buf := frame.GetRefVar(2)
	off := frame.GetIntVar(3)
	_len := frame.GetIntVar(4)

	conn := fd.Extra.(net.Conn)

	_timeout := frame.GetIntVar(5)
	if _timeout > 0 {
		conn.SetDeadline(time.Now().Add(time.Duration(_timeout) * time.Millisecond))
	}

	goBuf := buf.GetGoBytes()
	goBuf = goBuf[off : off+_len]

	n, err := conn.Read(goBuf)
	if err == nil || n > 0 || err == io.EOF {
		frame.PushInt(int32(n))
	} else {
		// todo
		panic(err.Error())
		//frame.Thread.ThrowIOException(err.Error())
	}

}
