package heap

type MethodDescriptor struct {
	d              string
	parameterTypes []*FieldType
	returnType     *FieldType
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{descriptor: descriptor}
	return parser.parse()
}

func (md *MethodDescriptor) String() string {
	return md.d
}

func (md *MethodDescriptor) ParameterTypes() []*FieldType {
	return md.parameterTypes
}
func (md *MethodDescriptor) ReturnType() *FieldType {
	return md.returnType
}

// parameterCount()
func (md *MethodDescriptor) argCount() uint {
	return uint(len(md.parameterTypes))
}

func (md *MethodDescriptor) argSlotCount() uint {
	slotCount := md.argCount()
	for _, paramType := range md.parameterTypes {
		if paramType.IsLongOrDouble() {
			slotCount++
		}
	}
	return slotCount
}

func (md *MethodDescriptor) addParameterType(t *FieldType) {
	pLen := len(md.parameterTypes)
	if pLen == cap(md.parameterTypes) {
		s := make([]*FieldType, pLen, pLen+4)
		copy(s, md.parameterTypes)
		md.parameterTypes = s
	}

	md.parameterTypes = append(md.parameterTypes, t)
}
