package heap

type FieldType string

func (ft FieldType) IsBaseType() bool {
	return len(ft) == 1
}
func (ft FieldType) IsVoidType() bool {
	return ft == "V"
}
func (ft FieldType) IsObjectType() bool {
	return ft[0] == 'L'
}
func (ft FieldType) IsArrayType() bool {
	return ft[0] == '['
}
func (ft FieldType) IsLongOrDouble() bool {
	return ft == "J" ||
		ft == "D"
}
