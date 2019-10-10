package heap

type FieldOrReturnType string

func (ft FieldOrReturnType) IsBaseType() bool     { return len(ft) == 1 }
func (ft FieldOrReturnType) IsVoidType() bool     { return ft == "V" }
func (ft FieldOrReturnType) IsObjectType() bool   { return ft[0] == 'L' }
func (ft FieldOrReturnType) IsArrayType() bool    { return ft[0] == '[' }
func (ft FieldOrReturnType) IsLongOrDouble() bool { return ft == "J" || ft == "D" }

type ParsedDescriptor struct {
	ParameterTypes []FieldOrReturnType
	ReturnType     FieldOrReturnType
}

func (md ParsedDescriptor) getParamCount() uint {
	return uint(len(md.ParameterTypes))
}

func (md ParsedDescriptor) getParamSlotCount() uint {
	slotCount := md.getParamCount()
	for _, paramType := range md.ParameterTypes {
		if paramType.IsLongOrDouble() {
			slotCount++
		}
	}
	return slotCount
}
