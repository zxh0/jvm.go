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
    AT_OBJEC     = 100 // not jvm spec
    AT_NOT_ARRAY = 101 // not jvm spec
)

func NewPrimitiveArray(atype uint8, count int32, classLoader *ClassLoader) (*Obj) {
    switch atype {
    case AT_BOOLEAN: return &Obj{classLoader.getClass("[Z"), make([]int8, count),      nil}
    case AT_BYTE:    return &Obj{classLoader.getClass("[B"), make([]int8, count),      nil}
    case AT_CHAR:    return &Obj{classLoader.getClass("[C"), make([]uint16, count),    nil}
    case AT_SHORT:   return &Obj{classLoader.getClass("[S"), make([]int16, count),     nil}
    case AT_INT:     return &Obj{classLoader.getClass("[I"), make([]int32, count),     nil}
    case AT_LONG:    return &Obj{classLoader.getClass("[J"), make([]int64, count),     nil}
    case AT_FLOAT:   return &Obj{classLoader.getClass("[F"), make([]float32, count),   nil}
    case AT_DOUBLE:  return &Obj{classLoader.getClass("[D"), make([]float64, count),   nil}
    default: panic("BAD atype!") // todo
    }
}

func NewRefArray(componentClass *Class, count int32, classLoader *ClassLoader) (*Obj) {
    arrClass := classLoader.getRefArrayClass(componentClass)
    components := make([]*Obj, count)
    return &Obj{arrClass, components, nil}
}

// todo only used by exec_main.go
func NewStringArray(strs []*Obj, classLoader *ClassLoader) (*Obj) {
    componentClass := classLoader.StringClass()
    arrClass := classLoader.getRefArrayClass(componentClass)
    return &Obj{arrClass, strs, nil}
}

// todo only used by string_helper.go
func NewIntArray(ints []int32) (*Obj) {
    return &Obj{nil, ints, nil}
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
