package class

const (
    //Array Type  atype
    AT_BOOLEAN   = 4
    AT_CHAR      = 5
    AT_FLOAT     = 6
    AT_DOUBLE    = 7
    AT_BYTE      = 8
    AT_SHORT     = 9
    AT_INT       = 10
    AT_LONG      = 11
)

func NewPrimitiveArray(atype uint8, count int32) (*Obj) {
    switch atype {
    case AT_BOOLEAN: return &Obj{make([]uint8, count)}
    case AT_CHAR:    return &Obj{make([]uint16, count)}
    case AT_FLOAT:   return &Obj{make([]float32, count)}
    case AT_DOUBLE:  return &Obj{make([]float64, count)}
    case AT_BYTE:    return &Obj{make([]uint8, count)}
    case AT_SHORT:   return &Obj{make([]int16, count)}
    case AT_INT:     return &Obj{make([]int32, count)}
    case AT_LONG:    return &Obj{make([]int64, count)}
    default: panic("BAD atype!") // todo
    }
}

func NewRefArray(count int32) (*Obj) {
    arr := make([]*Obj, count)
    return &Obj{arr}
}

func ArrayLength(arr *Obj) (int32) {
    // todo
    return 0
}
