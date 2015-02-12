package class

var _emptyParameterTypes = []*Class{}

func (self *Method) ParameterTypes() ([]*Class) {
    if self.argCount == 0 {
        return _emptyParameterTypes
    }

    md := self.MethodDescriptor()
    classLoader := self.class.classLoader
    parameterTypes := make([]*Class, self.argCount)
    for i, paramType := range md.parameterTypes {
        parameterTypes[i] = classLoader.LoadClass(paramType.descriptor) // todo
    }

    return parameterTypes
}
