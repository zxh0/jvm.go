package instructions

import (
    "jvmgo/jvm/rtda"
)

func branch(frame *rtda.Frame, offset int) {
    nextPC := frame.Thread().PC() + offset
    frame.SetNextPC(nextPC)
}

// todo
func checkArrIndex(index, len int) {
    if index < 0 || index >= len {
        panic("ArrayIndexOutOfBoundsException")
    }
}
