package class

func GetUpmostUninitializedClassOrInterface(from *Class) (*Class) {
    if from.initialized {
        return nil
    }
    loader := from.classLoader
    if from.superClassName != "" {
        superClass := loader.getClass(from.superClassName)
        if !superClass.initialized {
            return GetUpmostUninitializedClassOrInterface(superClass)
        }
    }
    for _, interfaceName := range from.interfaceNames {
        iClass := loader.getClass(interfaceName)
        if !iClass.initialized {
            return GetUpmostUninitializedClassOrInterface(iClass)
        }
    }
    return from
}
