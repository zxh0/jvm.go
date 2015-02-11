package class

// java primitive types
var primitiveTypes = map[string]string{
    "void":     "", 
    "boolean":  "[Z",
    "byte":     "[B", 
    "char":     "[C", 
    "short":    "[S", 
    "int":      "[I", 
    "long":     "[J", 
    "float":    "[F", 
    "double":   "[D",
}

func isPrimitiveType(name string) (bool) {
    _, ok := primitiveTypes[name]
    return ok
}
