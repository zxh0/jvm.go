package any

type Any interface{}

func IsLongOrDouble(x Any) (bool) {
    switch x.(type) {
    case int64: return true
    case float64: return true
    default: return false
    }
}
