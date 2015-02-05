package class

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

func (self FieldType) isBaseType() bool {
    return len(self.descriptor) == 1
}
func (self FieldType) isObjectType() bool {
    return self.descriptor[0] == 'L'
}
func (self FieldType) isArrayType() bool {
    return self.descriptor[0] == '['
}
