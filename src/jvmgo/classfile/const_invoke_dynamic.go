package classfile

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
    bootstrapMethodAttrIndex uint16
    nameAndTypeIndex         uint16
}

func readConstantInvokeDynamicInfo(reader *ClassReader) (*ConstantInvokeDynamicInfo) {
    bootstrapMethodAttrIndex := reader.readUint16()
    nameAndTypeIndex := reader.readUint16()
    return &ConstantInvokeDynamicInfo{bootstrapMethodAttrIndex, nameAndTypeIndex}
}
