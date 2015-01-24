package native

import (
    //"jvmgo/rtda"
    rtc "jvmgo/rtda/class"
)

// register native methods
func init() {
    rtc.RegisterNativeMethod("c","m","d", nil)
}
