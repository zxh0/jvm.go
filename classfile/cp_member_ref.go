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
type ConstantMemberRefInfo struct {
	Tag              uint8
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func readConstantMemberRefInfo(reader *ClassReader, tag uint8) ConstantMemberRefInfo {
	return ConstantMemberRefInfo{
		Tag:              tag,
		ClassIndex:       reader.readUint16(),
		NameAndTypeIndex: reader.readUint16(),
	}
}
