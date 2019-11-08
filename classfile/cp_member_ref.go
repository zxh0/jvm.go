package classfile

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/

type ConstantFieldRefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type ConstantMethodRefInfo struct {
	ClassIndex        uint16
	NameAndTypeIndex  uint16
	IsInterfaceMethod bool
}

func readConstantFieldRefInfo(reader *ClassReader) ConstantFieldRefInfo {
	return ConstantFieldRefInfo{
		ClassIndex:       reader.ReadUint16(),
		NameAndTypeIndex: reader.ReadUint16(),
	}
}

func readConstantMethodRefInfo(reader *ClassReader, isInterfaceMethod bool) ConstantMethodRefInfo {
	return ConstantMethodRefInfo{
		ClassIndex:        reader.ReadUint16(),
		NameAndTypeIndex:  reader.ReadUint16(),
		IsInterfaceMethod: isInterfaceMethod,
	}
}
