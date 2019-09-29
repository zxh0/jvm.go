package heap

type MethodDescriptor struct {
	d              string
	ParameterTypes []FieldType
	ReturnType     FieldType
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{descriptor: descriptor}
	return parser.parse()
}

func (md *MethodDescriptor) String() string {
	return md.d
}

// parameterCount()
func (md *MethodDescriptor) argCount() uint {
	return uint(len(md.ParameterTypes))
}

func (md *MethodDescriptor) argSlotCount() uint {
	slotCount := md.argCount()
	for _, paramType := range md.ParameterTypes {
		if paramType.IsLongOrDouble() {
			slotCount++
		}
	}
	return slotCount
}

func (md *MethodDescriptor) addParameterType(t FieldType) {
	pLen := len(md.ParameterTypes)
	if pLen == cap(md.ParameterTypes) {
		s := make([]FieldType, pLen, pLen+4)
		copy(s, md.ParameterTypes)
		md.ParameterTypes = s
	}

	md.ParameterTypes = append(md.ParameterTypes, t)
}
