package io

import (
	"fmt"
	"net"

	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_i6di(i6di_getHostByAddr, "getHostByAddr", "([B)Ljava/lang/String;")
	_i6di(i6di_lookupAllHostAddr, "lookupAllHostAddr", "(Ljava/lang/String;)[Ljava/net/InetAddress;")
}

func _i6di(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/net/Inet6AddressImpl", name, desc, method)
}

//([B)Ljava/lang/String;
//String getHostByAddr(byte[] var1) throws UnknownHostException
func i6di_getHostByAddr(frame *rtda.Frame) {
	vars := frame.LocalVars()
	t := vars.GetRef(1)
	buf := t.GoBytes()
	address := fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	if name, err := net.LookupAddr(address); err == nil {
		frame.OperandStack().PushRef(rtda.JString(name[0]))
	} else {
		panic("not found from lookupAllHostAddr")
		frame.OperandStack().PushRef(rtda.JString(""))
	}
}

//(Ljava/lang/String;)[Ljava/net/InetAddress;
func i6di_lookupAllHostAddr(frame *rtda.Frame) {
	vars := frame.LocalVars()
	host := rtda.GoString(vars.GetRef(1))
	address, _ := net.LookupHost(host)
	constructorCount := uint(len(address))

	inetAddress := heap.BootLoader().LoadClass("java/net/InetAddress")
	inetAddressArr := inetAddress.NewArray(constructorCount)

	stack := frame.OperandStack()
	stack.PushRef(inetAddressArr)

	//TODO
	//getByName descriptor:(Ljava/lang/String;)Ljava/net/InetAddress;
	//if constructorCount > 0 {
	//	thread := frame.Thread()
	//	constructorObjs := inetAddressArr.Refs()
	//	inetAddressGetByNameMethod := inetAddress.GetStaticMethod("getByName", "(Ljava/lang/String;)Ljava/net/InetAddress;")

	//	fmt.Println(constructorObjs[0])
	//	fmt.Println(inetAddressGetByNameMethod)
	//	fmt.Println(thread)
	//	thread.InvokeMethodWithShim(inetAddressGetByNameMethod, []interface{}{
	//		constructorObjs[0],
	//		rtda.JString(host),
	//	})
	//}
}
