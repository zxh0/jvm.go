package native

import (
    "jvmgo/jvm/rtda"
    //rtc "jvmgo/jvm/rtda/class"
)

type NativeMethod func(frame *rtda.Frame)
