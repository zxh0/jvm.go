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
    for _, i := range from.interfaces {
        if i.InitializationNotStarted() {
            return GetUpmostUninitializedClassOrInterface(i)
        }
    }
    return from
}
