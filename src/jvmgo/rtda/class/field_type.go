package class

import . "jvmgo/any"

func defaultValue(fieldDescriptor string) (Any) {
    switch fieldDescriptor[0] {
    case 'Z': // boolean
        fallthrough
    case 'B': // byte
        fallthrough
    case 'S': // short
        fallthrough
    case 'C': // char
        fallthrough
    case 'I': // int
        return int32(0)
    case 'J': // long
        return int64(0)
    case 'F': // float
        return float32(0)
    case 'D': // double
        return float64(0)
    case 'L': // Object
        return nil
    case '[': // Array
        return nil
    default: panic("BAD field descriptor: " + fieldDescriptor)
    }
}
