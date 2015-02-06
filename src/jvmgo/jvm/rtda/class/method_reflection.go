package class

func (self *Method) ParameterTypes() ([]*Class) {
    if self.argCount == 0 {
        // todo optimize
        return make([]*Class, 0)
    }

    // todo
    md := parseMethodDescriptor(self.descriptor)
    parameterTypes := make([]*Class, len(md.parameterTypes))
    for i, paramType := range md.parameterTypes {
        parameterTypes[i] = self.class.classLoader.LoadClass(paramType.descriptor) // todo
    }

    return parameterTypes
}
