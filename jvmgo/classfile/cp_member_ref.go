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
type ConstantMemberrefInfo struct {
	Tag              uint8
	cp               *ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}
func (self ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

func readConstantMemberrefInfo(reader *ClassReader,
	cp *ConstantPool, tag uint8) ConstantMemberrefInfo {

	return ConstantMemberrefInfo{
		Tag:              tag,
		cp:               cp,
		classIndex:       reader.readUint16(),
		nameAndTypeIndex: reader.readUint16(),
	}
}
