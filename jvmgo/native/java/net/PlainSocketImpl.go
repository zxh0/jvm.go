package io

import (
	"fmt"
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	"net"
)

func init() {
	_psi(psi_initProto, "initProto", "()V")
	_psi(psi_socketCreate, "socketCreate", "(Z)V")
	_psi(psi_socketBind, "socketBind", "(Ljava/net/InetAddress;I)V")
	_psi(psi_socketListen, "socketListen", "(I)V")
	_psi(psi_socketAccept, "socketAccept", "(Ljava/net/SocketImpl;)V")
	_psi(psi_socketClose0, "socketClose0", "(Z)V")
	_psi(psi_socketAvailable, "socketAvailable", "()I")
}

func _psi(method Any, name, desc string) {
	rtc.RegisterNativeMethod("java/net/PlainSocketImpl", name, desc, method)
}

func psi_initProto(frame *rtda.Frame) {
	//TODO
	//fmt.Println(frame.Method().Class().GetInstanceField("port", "I").Slot())

}

//  native void socketCreate(boolean isServer) throws IOException;
func psi_socketCreate(frame *rtda.Frame) {
	//fmt.Println(this.GetFieldValue("CONNECTION_RESET", "I").(int32))

}

//  native void socketBind(InetAddress address, int port)
func psi_socketBind(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	address := vars.GetRef(1)
	port := vars.Get(2)

	holder := address.GetFieldValue("holder", "Ljava/net/InetAddress$InetAddressHolder;").(*rtc.Obj)
	hostNameObj := holder.GetFieldValue("hostName", "Ljava/lang/String;").(*rtc.Obj)
	hostName := rtda.GoString(hostNameObj)
	laddr := fmt.Sprintf("%s:%d", hostName, port)
	listen, err := net.Listen("tcp", laddr)

	if err != nil {
		panic(err.Error())
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

	fdObj := s.GetFieldValue("fd", "Ljava/io/FileDescriptor;").(*rtc.Obj)
	//goFd := fdObj.GetFieldValue("fd", "I").(int32)
	///fmt.Printf("fd:%d\r\n", goFd)
	listen := this.Extra().(net.Listener)
	if conn, err := listen.Accept(); err != nil {
		panic(err.Error())
	} else {
		fdObj.SetExtra(conn)
		//fmt.Println("set connection!!!!")
		//conn.Write([]byte("Beyond\r\n"))
	}
}

// native void socketConnect(InetAddress address, int port, int timeout)
//        throws IOException;
// java/net/PlainSocketImpl~socketClose0~(Z)V
func psi_socketClose0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	fdObj := this.GetFieldValue("fd", "Ljava/io/FileDescriptor;").(*rtc.Obj)
	//goFd := fdObj.GetFieldValue("fd", "I").(int32)
	conn := fdObj.Extra().(net.Conn)
	conn.Close()
}

// native int socketAvailable() throws IOException;
// ()I
func psi_socketAvailable(frame *rtda.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	fdObj := this.GetFieldValue("fd", "Ljava/io/FileDescriptor;").(*rtc.Obj)
	conn := fdObj.Extra().(net.Conn)
	if r := conn.RemoteAddr(); r != nil {
		frame.OperandStack().PushBoolean(true)
	} else {
		frame.OperandStack().PushBoolean(false)
	}

}
