package classfile

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfo struct {
    descriptorIndex uint16
}

func readConstantMethodTypeInfo(reader *ClassReader) (*ConstantMethodTypeInfo) {
    descriptorIndex := reader.readUint16()
    return &ConstantMethodTypeInfo{descriptorIndex}
}
