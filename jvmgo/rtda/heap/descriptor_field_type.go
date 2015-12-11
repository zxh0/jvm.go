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

func (self *FieldType) Descriptor() string {
	return self.descriptor
}

func (self *FieldType) IsBaseType() bool {
	return len(self.descriptor) == 1
}
func (self *FieldType) IsVoidType() bool {
	return self.descriptor == "V"
}
func (self *FieldType) IsObjectType() bool {
	return self.descriptor[0] == 'L'
}
func (self *FieldType) IsArrayType() bool {
	return self.descriptor[0] == '['
}
func (self *FieldType) IsLongOrDouble() bool {
	return self.descriptor == "J" ||
		self.descriptor == "D"
}
