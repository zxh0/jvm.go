package misc

import (
	"math"
	"time"

	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

func init() {
	_unsafe(park, "park", "(ZJ)V")
	_unsafe(unpark, "unpark", "(Ljava/lang/Object;)V")
}

func _unsafe(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("sun/misc/Unsafe", name, desc, method)
}

// public native void park(boolean isAbsolute, long time);
// (ZJ)V [
/*
http://www.docjar.com/docs/api/sun/misc/Unsafe.html#park%28boolean,%20long%29
Block current thread, returning when a balancing
unpark occurs, or a balancing unpark has
already occurred, or the thread is interrupted, or, if not
absolute and time is not zero, the given time nanoseconds have
elapsed, or if absolute, the given deadline in milliseconds
since Epoch has passed, or spuriously (i.e., returning for no
"reason"). Note: This operation is in the Unsafe class only
because unpark is, so it would be strange to place it
elsewhere.
*/
func park(frame *rtda.Frame) {
	vars := frame.LocalVars()
	isAbsolute := vars.GetBoolean(1)
	var2 := vars.GetLong(2)
	var parkTime time.Duration

	if isAbsolute {
		deadline := var2 * int64(time.Millisecond)
		now := time.Now().UnixNano()
		parkTime = time.Duration(deadline - now)
	} else {
		if var2 == 0 { // park forever
			parkTime = time.Duration(math.MaxInt64)
		} else {
			parkTime = time.Duration(var2)
		}
	}

	frame.Thread().Park(parkTime)
}

//  public native void unpark(Object thread);
//  (Ljava/lang/Object;)V
func unpark(frame *rtda.Frame) {
	vars := frame.LocalVars()
	threadObj := vars.GetRef(1)

	thread := threadObj.Extra().(*rtda.Thread)
	thread.Unpark()
}
