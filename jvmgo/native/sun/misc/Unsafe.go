package misc

import (
	. "github.com/zxh0/jvm.go/jvmgo/any"
	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	"time"
)

func init() {
	_unsafe(park, "park", "(ZJ)V")
	_unsafe(unpark, "unpark", "(Ljava/lang/Object;)V")
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
		parkTime = time.Duration((var2 - (time.Now().UnixNano() / 1000000)) * int64(time.Millisecond))
	} else {
		parkTime = time.Duration(var2)
	}
	frame.Thread().Park(parkTime)
}

//  public native void unpark(Object var1);
//  (Ljava/lang/Object;)V
func unpark(frame *rtda.Frame) {
	vars := frame.LocalVars()
	thread := vars.GetRef(1).Extra().(*rtda.Thread)
	thread.Unpark()
}
