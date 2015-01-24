package native

import (
    "jvmgo/rtda"
    //rtc "jvmgo/rtda/class"
)

type NativeMethod func(operandStack *rtda.OperandStack)
