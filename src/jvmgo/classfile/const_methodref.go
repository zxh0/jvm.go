package classfile

/*
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantMethodrefInfo struct {
    classIndex       uint16
    nameAndTypeIndex uint16
}

func readConstantMethodrefInfo(reader *ClassReader) (*ConstantMethodrefInfo) {
    classIndex := reader.readUint16()
    nameAndTypeIndex := reader.readUint16()
    return &ConstantMethodrefInfo{classIndex, nameAndTypeIndex}
}

/*
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantInterfaceMethodrefInfo struct {
    classIndex       uint16
    nameAndTypeIndex uint16
}

func readConstantInterfaceMethodrefInfo(reader *ClassReader) (*ConstantInterfaceMethodrefInfo) {
    classIndex := reader.readUint16()
    nameAndTypeIndex := reader.readUint16()
    return &ConstantInterfaceMethodrefInfo{classIndex, nameAndTypeIndex}
}
