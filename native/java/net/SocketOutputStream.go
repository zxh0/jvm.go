package io

import (
	"net"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_sos(sos_init, "init", "()V")
	_sos(sos_socketWrite0, "socketWrite0", "(Ljava/io/FileDescriptor;[BII)V")
}

func _sos(method native.Method, name, desc string) {
	native.Register("java/net/SocketOutputStream", name, desc, method)
}

// native java/net/SocketOutputStream~init~()V
func sos_init(frame *rtda.Frame) {

}

//    private native void socketWrite0(FileDescriptor fd, byte[] b, int off,
//                                     int len) throws IOException;
// java/net/SocketOutputStream~socketWrite0~(Ljava/io/FileDescriptor;
func sos_socketWrite0(frame *rtda.Frame) {
	//this := frame.GetThis()
	fd := frame.GetRefVar(1)
	b := frame.GetRefVar(2)
	offset := frame.GetIntVar(3)
	length := frame.GetIntVar(4)
	conn := fd.Extra.(net.Conn)
	goBytes := b.GetGoBytes()
	goBytes = goBytes[offset : offset+length]

	if _, err := conn.Write(goBytes); err != nil {
		frame.Thread.ThrowIOException(err.Error())
	}
}
