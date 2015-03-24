package io

import (
	"fmt"
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	"github.com/zxh0/jvm.go/jvmgo/util"
	"net"
)

func init() {
	_sos(sos_init, "init", "()V")
	_sos(sos_socketWrite0, "socketWrite0", "(Ljava/io/FileDescriptor;[BII)V")
}

func _sos(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/net/SocketOutputStream", name, desc, method)
}

// native java/net/SocketOutputStream~init~()V
func sos_init(frame *rtda.Frame) {

}

//    private native void socketWrite0(FileDescriptor fd, byte[] b, int off,
//                                     int len) throws IOException;
// java/net/SocketOutputStream~socketWrite0~(Ljava/io/FileDescriptor;
func sos_socketWrite0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	//this := vars.GetThis()
	fd := vars.GetRef(1)
	b := vars.GetRef(2)
	offset := vars.GetInt(3)
	length := vars.GetInt(4)

	conn := fd.Extra().(net.Conn)
	fmt.Println("########################")

	jBytes := b.Fields().([]int8)
	jBytes = jBytes[offset : offset+length]
	goBytes := util.CastInt8sToUint8s(jBytes)
	conn.Write(goBytes)
}
