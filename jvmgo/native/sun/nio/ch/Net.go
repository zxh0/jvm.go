//sun/nio/ch/ServerSocketChannelImpl~initIDs~
package ch

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
)

func init() {
	_net(net_isExclusiveBindAvailable, "isExclusiveBindAvailable", "()I")
	//_net(net_pollinValue, "pollinValue", "()S")
}

func _net(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/nio/ch/Net", name, desc, method)
}

func net_isExclusiveBindAvailable(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

func net_pollinValue(frame *rtda.Frame) {

}
