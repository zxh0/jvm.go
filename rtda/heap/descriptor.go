package heap

type TypeDescriptor string

func (td TypeDescriptor) IsBaseType() bool     { return len(td) == 1 }
func (td TypeDescriptor) IsVoidType() bool     { return td == "V" }
func (td TypeDescriptor) IsObjectType() bool   { return td[0] == 'L' }
func (td TypeDescriptor) IsArrayType() bool    { return td[0] == '[' }
func (td TypeDescriptor) IsLongOrDouble() bool { return td == "J" || td == "D" }

type MethodDescriptor struct {
	ParameterTypes []TypeDescriptor
	ReturnType     TypeDescriptor
}

func (md MethodDescriptor) getParamCount() uint {
	return uint(len(md.ParameterTypes))
}

func (md MethodDescriptor) getParamSlotCount() uint {
	slotCount := md.getParamCount()
	for _, paramType := range md.ParameterTypes {
		if paramType.IsLongOrDouble() {
			slotCount++
		}
	}
	return slotCount
}
