package classfile

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantFieldrefInfo struct {
    classIndex       uint16
    nameAndTypeIndex uint16
}

func readConstantFieldrefInfo(reader *ClassReader) (*ConstantFieldrefInfo) {
    classIndex := reader.readUint16()
    nameAndTypeIndex := reader.readUint16()
    return &ConstantFieldrefInfo{classIndex, nameAndTypeIndex}
}
