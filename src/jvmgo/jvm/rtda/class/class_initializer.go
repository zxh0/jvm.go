package class

func GetUpmostUninitializedClassOrInterface(from *Class) (*Class) {
    if !from.InitializationNotStarted() {
        return nil
    }
    for k := from.superClass; k != nil; k = k.superClass {
        if k.InitializationNotStarted() {
            return GetUpmostUninitializedClassOrInterface(k)
        }
    }

    loader := from.classLoader
    for _, interfaceName := range from.interfaceNames {
        iClass := loader.getClass(interfaceName)
        if iClass.InitializationNotStarted() {
            return GetUpmostUninitializedClassOrInterface(iClass)
        }
    }
    return from
}
