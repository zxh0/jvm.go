package instructions

import (
    . "jvmgo/any"
    "jvmgo/rtda"
)

func branch(thread *rtda.Thread, offset int) {
    nextPC := thread.PC() + offset
    thread.CurrentFrame().SetNextPC(nextPC)
}

// todo: move to any.go?
func isLongOrDouble(x Any) (bool) {
    switch x.(type) {
    case int64: return true
    case float64: return true
    default: return false
    }
}

// todo
func checkArrIndex(index, len int) {
    if index < 0 || index >= len {
        panic("ArrayIndexOutOfBoundsException")
    }
}
