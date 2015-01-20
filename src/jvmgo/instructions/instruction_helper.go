package instructions

import (
    . "jvmgo/any"
    //"jvmgo/rtda"
)

// todo: move to any.go?
func isLongOrDouble(x Any) (bool) {
    switch x.(type) {
    case int64: return true
    case float64: return true
    default: return false
    }
}
