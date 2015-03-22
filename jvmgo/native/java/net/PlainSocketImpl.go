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
	listen := this.Extra().(net.Listener)
	if conn, err := listen.Accept(); err != nil {
		panic(err.Error())
	} else {
		conn.Write([]byte("Beyond\r\n"))
	}
}

// native void socketConnect(InetAddress address, int port, int timeout)
//        throws IOException;
