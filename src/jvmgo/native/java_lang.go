package native

import (
    "time"
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// register native methods
func init() {
    rtc.RegisterNativeMethod("java/lang/System","nanoTime","()J", nanoTime)
    rtc.RegisterNativeMethod("java/lang/System","currentTimeMillis","()J", nanoTime)
}

// java.lang.System
func nanoTime(stack *rtda.OperandStack) {
    nanoTime := time.Now().UnixNano()
    stack.PushLong(nanoTime)
}
func currentTimeMillis(stack *rtda.OperandStack) {
    millis := time.Now().UnixNano() / 1000
    stack.PushLong(millis)
}
