package native

import (
    "jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

var system_nanoTime = func(stack *rtda.OperandStack) {
    // todo
    panic("nannannanatime...")
}

// register native methods
func init() {
    rtc.RegisterNativeMethod("java/lang/System","nanoTime","()J", system_nanoTime)
}
