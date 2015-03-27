package misc

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	"time"
)

func init() {
	_unsafe(park, "park", "(ZJ)V")
}

func _unsafe(method Any, name, desc string) {
	rtc.RegisterNativeMethod("sun/misc/Unsafe", name, desc, method)
}


// public native void park(boolean var1, long var2);
// (ZJ)V [
func park(frame *rtda.Frame) {
	vars := frame.LocalVars()
	absolute := vars.GetBoolean(1)
	var2 := vars.GetLong(2)
	var parkTime time.Duration

	//deadline the absolute time, in milliseconds from the Epoch,
	// *        to wait until
	if absolute {
		parkTime = time.Duration((time.Now().UnixNano() / 1000) - var2)
	} else {
		parkTime = time.Duration(var2)
	}
	frame.Thread().Park(parkTime)
}
