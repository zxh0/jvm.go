package ch

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/zxh0/jvm.go/native"
	"github.com/zxh0/jvm.go/rtda"
)

func init() {
	_net(net_isExclusiveBindAvailable, "isExclusiveBindAvailable", "()I")
	_net(net_pollinValue, "pollinValue", "()S")
	_net(net_pollinValue, "polloutValue", "()S")
	_net(net_pollinValue, "pollerrValue", "()S")
	_net(net_pollinValue, "pollhupValue", "()S")
	_net(net_pollinValue, "pollnvalValue", "()S")
	_net(net_pollinValue, "pollconnValue", "()S")
	_net(net_isIPv6Available, "isIPv6Available0", "()Z")

	_net(net_socket0, "socket0", "(ZZZ)I")
	_net(net_setIntOption0, "setIntOption0", "(Ljava/io/FileDescriptor;ZIII)V")
	_net(net_bind0, "bind0", "(Ljava/io/FileDescriptor;ZZLjava/net/InetAddress;I)V")
	_net(net_listen, "listen", "(Ljava/io/FileDescriptor;I)V")
	_net(net_localInetAddress, "localInetAddress", "(Ljava/io/FileDescriptor;)Ljava/net/InetAddress;")
	_net(net_localPort, "localPort", "(Ljava/io/FileDescriptor;)I")
}

func _net(method native.Method, name, desc string) {
	native.Register("sun/nio/ch/Net", name, desc, method)
}

func net_isExclusiveBindAvailable(frame *rtda.Frame) {
	frame.PushInt(-1)
}

func net_pollinValue(frame *rtda.Frame) {
	frame.PushInt(1)
}

// Tells whether dual-IPv4/IPv6 sockets should be used.
func net_isIPv6Available(frame *rtda.Frame) {
	frame.PushBoolean(true)
}

// Due to oddities SO_REUSEADDR on windows reuse is ignored
// private static native int socket0(boolean preferIPv6, boolean stream, boolean reuse);
func net_socket0(frame *rtda.Frame) {
	//vars := frame.
	//preferIPv6 := frame.GetBooleanVar(0)
	//stream := frame.GetBooleanVar(1)
	//reuse := frame.GetBooleanVar(2)

	frame.PushInt(100)
}

// private static native void setIntOption0(FileDescriptor fd, boolean mayNeedConversion, int level, int opt, int arg)
//(Ljava/io/FileDescriptor;ZIII)V
func net_setIntOption0(frame *rtda.Frame) {

}

// private static native void bind0(FileDescriptor fd, boolean preferIPv6, boolean useExclBind, InetAddress addr,int port)
func net_bind0(frame *rtda.Frame) {
	this := frame.GetThis()
	address := frame.GetRefVar(3)
	port := frame.GetLocalVar(4)

	holder := address.GetFieldValue("holder", "Ljava/net/InetAddress$InetAddressHolder;").Ref
	hostNameObj := holder.GetFieldValue("hostName", "Ljava/lang/String;").Ref
	hostName := hostNameObj.JSToGoStr()
	laddr := fmt.Sprintf("%s:%d", hostName, port.IntValue())
	listen, err := net.Listen("tcp", laddr)
	if err != nil {
		frame.Thread.ThrowIOException(err.Error())
	}
	this.Extra = listen
}

// static native void listen(FileDescriptor fd, int backlog) throws IOException;
func net_listen(frame *rtda.Frame) {

}

// private static native InetAddress localInetAddress(FileDescriptor fd) throws IOException;
func net_localInetAddress(frame *rtda.Frame) {
	//vars := frame.
	//this := frame.GetThis()
	//listen := this.Extra().(net.Listener)

	inetAddress := frame.GetBootLoader().LoadClass("java/net/InetAddress")
	inetObj := inetAddress.NewObj()

	frame.PushRef(inetObj)

	//fmt.Println(inetAddress)
	//fmt.Println(listen.Addr().String())
	//fmt.Println(listen.Addr().Network())

	//panic("net_localInetAddress error")
}

func net_localPort(frame *rtda.Frame) {
	this := frame.GetThis()
	listen := this.Extra.(net.Listener)

	//fmt.Println(inetAddress)
	address := listen.Addr().String()
	lastIndex := strings.LastIndex(address, ":")

	port, _ := strconv.Atoi(address[lastIndex+1:])
	frame.PushInt(int32(port))
}
