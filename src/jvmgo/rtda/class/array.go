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
    case AT_BOOLEAN: return &Obj{nil, make([]int8, count)}
    case AT_BYTE:    return &Obj{nil, make([]int8, count)}
    case AT_CHAR:    return &Obj{nil, make([]uint16, count)}
    case AT_SHORT:   return &Obj{nil, make([]int16, count)}
    case AT_INT:     return &Obj{nil, make([]int32, count)}
    case AT_LONG:    return &Obj{nil, make([]int64, count)}
    case AT_FLOAT:   return &Obj{nil, make([]float32, count)}
    case AT_DOUBLE:  return &Obj{nil, make([]float64, count)}
    default: panic("BAD atype!") // todo
    }
}

func NewRefArray(count int32) (*Obj) {
    elements := make([]*Obj, count)
    return &Obj{nil, elements}
}
func NewRefArrayOfElements(elements []*Obj) (*Obj) {
    return &Obj{nil, elements}
}

func NewIntArray(ints []int32) (*Obj) {
    return &Obj{nil, ints}
}

// todo
func IsArray(obj *Obj) (bool) {
    return obj.class == nil
}

func ArrayLength(arr *Obj) (int32) {
    switch arr.fields.(type) {
        case []int8: return int32(len(arr.fields.([]int8)))
        case []int16: return int32(len(arr.fields.([]int16)))
        case []int32: return int32(len(arr.fields.([]int32)))
        case []int64: return int32(len(arr.fields.([]int64)))
        case []uint16: return int32(len(arr.fields.([]uint16)))
        case []float32: return int32(len(arr.fields.([]float32)))
        case []float64: return int32(len(arr.fields.([]float64)))
        case []*Obj: return int32(len(arr.fields.([]*Obj)))
        default: panic("Not array!") // todo
    }
}

func ArrayCopy(src, dest *Obj, srcPos, destPos, length int32) {



    // todo
    panic("ArrayCopy")
}
