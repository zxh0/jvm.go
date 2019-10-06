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

type ConstantFieldRefInfo constantMemberRefInfo
type ConstantMethodRefInfo constantMemberRefInfo
type ConstantInterfaceMethodRefInfo constantMemberRefInfo

type constantMemberRefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func readConstantFieldRefInfo(reader *ClassReader) ConstantFieldRefInfo {
	return ConstantFieldRefInfo(readConstantMemberRefInfo(reader))
}
func readConstantMethodRefInfo(reader *ClassReader) ConstantMethodRefInfo {
	return ConstantMethodRefInfo(readConstantMemberRefInfo(reader))
}
func readConstantInterfaceMethodRefInfo(reader *ClassReader) ConstantInterfaceMethodRefInfo {
	return ConstantInterfaceMethodRefInfo(readConstantMemberRefInfo(reader))
}

func readConstantMemberRefInfo(reader *ClassReader) constantMemberRefInfo {
	return constantMemberRefInfo{
		ClassIndex:       reader.ReadUint16(),
		NameAndTypeIndex: reader.ReadUint16(),
	}
}
