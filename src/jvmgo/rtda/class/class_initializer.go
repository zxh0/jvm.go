package class

func GetUpmostUninitializedClassOrInterface(from *Class) (*Class) {
    if !from.InitializationNotStarted() {
        return nil
    }
    loader := from.classLoader
    if from.superClassName != "" {
        superClass := loader.getClass(from.superClassName)
        if superClass.InitializationNotStarted() {
            return GetUpmostUninitializedClassOrInterface(superClass)
        }
    }
    for _, interfaceName := range from.interfaceNames {
        iClass := loader.getClass(interfaceName)
        if iClass.InitializationNotStarted() {
            return GetUpmostUninitializedClassOrInterface(iClass)
        }
    }
    return from
}
