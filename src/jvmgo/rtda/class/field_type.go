package class

import . "jvmgo/any"

func defaultValue(fieldDescriptor string) (Any) {
    switch fieldDescriptor[0] {
    case 'Z': return int32(0)   // boolean
    case 'B': return int32(0)   // byte
    case 'S': return int32(0)   // short
    case 'C': return int32(0)   // char
    case 'I': return int32(0)   // int
    case 'J': return int64(0)   // long
    case 'F': return float32(0) // float
    case 'D': return float64(0) // double
    case 'L': return nil        // Object
    case '[': return nil        // Array
    default: panic("BAD field descriptor: " + fieldDescriptor)
    }
}
