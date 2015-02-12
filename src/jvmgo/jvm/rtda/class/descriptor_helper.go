package class

import "strings"

func calcArgCount(descriptor string) (uint) {
    return parseMethodDescriptor(descriptor).argCount()
}

// func descriptorToClassName(descriptor string) string {
//     switch descriptor[0] {
//     case '[': return descriptor // array
//     case 'L': return descriptor[1, len(descriptor) - 1] // object
//     }
// }

func GetReturnDescriptor(methodDescriptor string) string {
    start := strings.Index(methodDescriptor, ")") + 1
    return methodDescriptor[start:]
}
