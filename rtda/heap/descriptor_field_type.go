package heap

var (
	baseTypeB = &FieldType{"B"} // byte
	baseTypeC = &FieldType{"C"} // char
	baseTypeD = &FieldType{"D"} // double
	baseTypeF = &FieldType{"F"} // float
	baseTypeI = &FieldType{"I"} // int
	baseTypeJ = &FieldType{"J"} // long
	baseTypeS = &FieldType{"S"} // short
	baseTypeZ = &FieldType{"Z"} // boolean
	baseTypeV = &FieldType{"V"} // void
)

type FieldType struct {
	descriptor string
}

func (ft *FieldType) Descriptor() string {
	return ft.descriptor
}

func (ft *FieldType) IsBaseType() bool {
	return len(ft.descriptor) == 1
}
func (ft *FieldType) IsVoidType() bool {
	return ft.descriptor == "V"
}
func (ft *FieldType) IsObjectType() bool {
	return ft.descriptor[0] == 'L'
}
func (ft *FieldType) IsArrayType() bool {
	return ft.descriptor[0] == '['
}
func (ft *FieldType) IsLongOrDouble() bool {
	return ft.descriptor == "J" ||
		ft.descriptor == "D"
}
