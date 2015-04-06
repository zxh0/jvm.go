package class

type MethodDescriptor struct {
	d              string
	parameterTypes []*FieldType
	returnType     *FieldType
}

func (self *MethodDescriptor) String() string {
	return self.d
}

func (self *MethodDescriptor) ParameterTypes() []*FieldType {
	return self.parameterTypes
}
func (self *MethodDescriptor) ReturnType() *FieldType {
	return self.returnType
}

// parameterCount()
func (self *MethodDescriptor) argCount() uint {
	return uint(len(self.parameterTypes))
}

func (self *MethodDescriptor) addParameterType(t *FieldType) {
	pLen := len(self.parameterTypes)
	if pLen == cap(self.parameterTypes) {
		s := make([]*FieldType, pLen, pLen+4)
		copy(s, self.parameterTypes)
		self.parameterTypes = s
	}

	self.parameterTypes = append(self.parameterTypes, t)
}
