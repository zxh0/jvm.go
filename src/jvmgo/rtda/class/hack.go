package class

// only used by exec_main.go
func NewStringArray(strs []*Obj, classLoader *ClassLoader) (*Obj) {
    componentClass := classLoader.StringClass()
    arrClass := classLoader.getRefArrayClass(componentClass)
    return &Obj{arrClass, strs, nil}
}

// only used by string_helper.go
func NewIntArray(ints []int32) (*Obj) {
    return &Obj{nil, ints, nil}
}
