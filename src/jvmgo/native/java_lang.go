package native

import (
    "time"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// register native methods
func init() {
    rtc.RegisterNativeMethod("java/lang/System","nanoTime","()J", system_nanoTime)
}

func system_nanoTime(stack *rtda.OperandStack) {
    nanoTime := time.Now().UnixNano()
    stack.PushLong(nanoTime)
}
