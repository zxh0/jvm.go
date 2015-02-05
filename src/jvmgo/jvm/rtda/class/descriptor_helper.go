package class

import (
    "strings"
)

func calcArgCount(descriptor string) (uint) {
    return parseMethodDescriptor(descriptor).argCount()
}

func isVoidReturnType(descriptor string) bool {
    return strings.HasSuffix(descriptor, ")V")
}
