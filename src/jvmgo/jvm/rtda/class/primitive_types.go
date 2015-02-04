package class

// java primitive types
var primitiveTypes = []string{
    "void", "boolean", "byte", "char", "short", "int", "long", "float", "double",
}

func isPrimitiveType(name string) (bool) {
    for _, primitiveType := range primitiveTypes {
        if primitiveType == name {
            return true
        }
    }
    return false
}
