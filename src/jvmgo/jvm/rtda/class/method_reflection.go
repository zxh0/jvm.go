package class

func (self *Method) ParameterTypes() ([]*Class) {
    if self.argCount == 0 {
        // todo optimize
        return make([]*Class, 0)
    }

    md := self.MethodDescriptor()
    parameterTypes := make([]*Class, self.argCount)
    for i, paramType := range md.parameterTypes {
        parameterTypes[i] = self.class.classLoader.LoadClass(paramType.descriptor) // todo
    }

    return parameterTypes
}
