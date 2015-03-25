package io

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	"io"
	"net"
)

func init() {
	_sis(sis_init, "init", "()V")
	_sis(sis_socketRead0, "socketRead0", "(Ljava/io/FileDescriptor;[BIII)I")
}

func _sis(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/net/SocketInputStream", name, desc, method)
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

	//TODO what ? timeout how to implement ?
	//_timeout := vars.GetInt(5)
	//conn.SetReadDeadline()

	goBuf := buf.GoBytes()
	goBuf = goBuf[off : off+_len]

	n, err := conn.Read(goBuf)
	//fmt.Println(string(goBuf))
	if err == nil || n > 0 || err == io.EOF {
		frame.OperandStack().PushInt(int32(n))
	} else {
		// todo
		panic(err.Error())
		//frame.Thread().ThrowIOException(err.Error())
	}

}
