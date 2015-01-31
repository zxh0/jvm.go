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
    AT_OBJEC     = 100 // no jvm spec
    AT_NOT_ARRAY = 101 // no jvm spec
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
func HaveSameArrayType(obj1, obj2 *Obj) (bool) {
    at1 := _arrayType(obj1)
    if at1 == AT_NOT_ARRAY {
        return false
    }

    at2 := _arrayType(obj2)
    if at2 == AT_NOT_ARRAY {
        return false
    }

    return at1 == at2
}

func _arrayType(arr *Obj) (int) {
    switch arr.fields.(type) {
        case []int8: return AT_BYTE
        case []int16: return AT_SHORT
        case []int32: return AT_INT
        case []int64: return AT_LONG
        case []uint16: return AT_CHAR
        case []float32: return AT_FLOAT
        case []float64: return AT_DOUBLE
        case []*Obj: return AT_OBJEC
        default: return AT_NOT_ARRAY
    }
}

// todo GetArrayLength
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

func ArrayCopy(src, dst *Obj, srcPos, dstPos, length int32) {
    switch src.fields.(type) {
        case []int8: 
            _src := src.fields.([]int8)[srcPos:srcPos+length]
            _dst := dst.fields.([]int8)[dstPos:dstPos+length]
            copy(_dst, _src)
        case []int16:
            _src := src.fields.([]int16)[srcPos:srcPos+length]
            _dst := dst.fields.([]int16)[dstPos:dstPos+length]
            copy(_dst, _src)
        case []int32:
            _src := src.fields.([]int32)[srcPos:srcPos+length]
            _dst := dst.fields.([]int32)[dstPos:dstPos+length]
            copy(_dst, _src)
        case []int64:
            _src := src.fields.([]int64)[srcPos:srcPos+length]
            _dst := dst.fields.([]int64)[dstPos:dstPos+length]
            copy(_dst, _src)
        case []uint16:
            _src := src.fields.([]uint16)[srcPos:srcPos+length]
            _dst := dst.fields.([]uint16)[dstPos:dstPos+length]
            copy(_dst, _src)
        case []float32:
            _src := src.fields.([]float32)[srcPos:srcPos+length]
            _dst := dst.fields.([]float32)[dstPos:dstPos+length]
            copy(_dst, _src)
        case []float64:
            _src := src.fields.([]float64)[srcPos:srcPos+length]
            _dst := dst.fields.([]float64)[dstPos:dstPos+length]
            copy(_dst, _src)
        case []*Obj:
            _src := src.fields.([]*Obj)[srcPos:srcPos+length]
            _dst := dst.fields.([]*Obj)[dstPos:dstPos+length]
            copy(_dst, _src)
        default: panic("Not array!") // todo
    }
}
