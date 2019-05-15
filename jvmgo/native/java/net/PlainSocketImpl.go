package io

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/zxh0/jvm.go/jvmgo/native/box"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_psi(psi_initProto, "initProto", "()V")
	_psi(psi_socketCreate, "socketCreate", "(Z)V")
	_psi(psi_socketBind, "socketBind", "(Ljava/net/InetAddress;I)V")
	_psi(psi_socketListen, "socketListen", "(I)V")
	_psi(psi_socketAccept, "socketAccept", "(Ljava/net/SocketImpl;)V")
	_psi(psi_socketClose0, "socketClose0", "(Z)V")
	_psi(psi_socketAvailable, "socketAvailable", "()I")
	_psi(psi_socketConnect, "socketConnect", "(Ljava/net/InetAddress;II)V")
	_psi(psi_socketSetOption, "socketSetOption", "(IZLjava/lang/Object;)V")

}

func _psi(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/net/PlainSocketImpl", name, desc, method)
}

func psi_initProto(frame *rtda.Frame) {
	//TODO

}

//  native void socketCreate(boolean isServer) throws IOException;
func psi_socketCreate(frame *rtda.Frame) {
	//fmt.Println(this.GetFieldValue("CONNECTION_RESET", "I").(int32))
}

//  native void socketConnect(InetAddress address, int port, int timeout)
//  throws IOException;
func psi_socketConnect(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	address := vars.GetRef(1)
	port := vars.Get(2)

	holder := address.GetFieldValue("holder", "Ljava/net/InetAddress$InetAddressHolder;").(*heap.Object)

	//fmt.Println(address.Class().GetInstanceMethod("getHostAddress", "()Ljava/lang/String;").NativeMethod())
	add := holder.GetFieldValue("address", "I").(int32)
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(add))
	laddr := fmt.Sprintf("%d.%d.%d.%d:%d", b[0], b[1], b[2], b[3], port)

	conn, err := net.Dial("tcp", laddr)

	if err != nil {
		frame.Thread().ThrowIOException(err.Error())
	}
	//TODO what ? timeout how to implement ?
	_timeout := vars.GetInt(3)
	if _timeout > 0 {
		conn.SetDeadline(time.Now().Add(time.Duration(_timeout) * time.Millisecond))
	}

	fdObj := this.GetFieldValue("fd", "Ljava/io/FileDescriptor;").(*heap.Object)
	fdObj.SetExtra(conn)
}

//  native void socketBind(InetAddress address, int port)
//	throws IOException
func psi_socketBind(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	address := vars.GetRef(1)
	port := vars.Get(2)

	holder := address.GetFieldValue("holder", "Ljava/net/InetAddress$InetAddressHolder;").(*heap.Object)
	hostNameObj := holder.GetFieldValue("hostName", "Ljava/lang/String;").(*heap.Object)
	hostName := rtda.GoString(hostNameObj)
	laddr := fmt.Sprintf("%s:%d", hostName, port)
	listen, err := net.Listen("tcp", laddr)

	if err != nil {
		frame.Thread().ThrowIOException(err.Error())
	}
	this.SetExtra(listen)
}

//  native void socketListen(int count) throws IOException;
func psi_socketListen(frame *rtda.Frame) {

}

//  native void socketAccept(SocketImpl s)
func psi_socketAccept(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	s := vars.GetRef(1)
	fdObj := s.GetFieldValue("fd", "Ljava/io/FileDescriptor;").(*heap.Object)
	//goFd := fdObj.GetFieldValue("fd", "I").(int32)
	listen := this.Extra().(net.Listener)
	if conn, err := listen.Accept(); err != nil {
		frame.Thread().ThrowIOException(err.Error())
	} else {
		fdObj.SetExtra(conn)
	}
}

// native void socketConnect(InetAddress address, int port, int timeout)
//        throws IOException;
// java/net/PlainSocketImpl~socketClose0~(Z)V
func psi_socketClose0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	fdObj := this.GetFieldValue("fd", "Ljava/io/FileDescriptor;").(*heap.Object)
	if fdObj.Extra() != nil {
		conn := fdObj.Extra().(net.Conn)
		conn.Close()
	}
}

// native int socketAvailable() throws IOException;
// ()I
func psi_socketAvailable(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	fdObj := this.GetFieldValue("fd", "Ljava/io/FileDescriptor;").(*heap.Object)
	conn := fdObj.Extra().(net.Conn)
	if r := conn.RemoteAddr(); r != nil {
		frame.OperandStack().PushBoolean(true)
	} else {
		frame.OperandStack().PushBoolean(false)
	}

}

// java/net/PlainSocketImpl~socketSetOption~
// (IZLjava/lang/Object;)V
// abstract void socketSetOption(int cmd, boolean on, Object value)
//  throws SocketException;
func psi_socketSetOption(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	cmd := vars.GetInt(1)
	//on := vars.GetBoolean(2)
	value := vars.GetRef(3)

	switch cmd {
	case 0x1006: //timeout
		_timeout := box.Unbox(value, "I").(int32)
		fdObj := this.GetFieldValue("fd", "Ljava/io/FileDescriptor;").(*heap.Object)
		if fdObj.Extra() != nil {
			conn := fdObj.Extra().(net.Conn)
			conn.SetReadDeadline(time.Now().Add(time.Duration(_timeout) * time.Millisecond))
		}
		break
	default:
		break
	}

}
