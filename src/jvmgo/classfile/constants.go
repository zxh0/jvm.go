package classfile

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/
type ConstantNameAndTypeInfo struct {
    nameIndex       uint16
    descriptorIndex uint16
}
func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
    self.nameIndex = reader.readUint16()
    self.descriptorIndex = reader.readUint16()
}

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
    referenceKind  uint8
    referenceIndex uint16
}
func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
    self.referenceKind = reader.readUint8()
    self.referenceIndex = reader.readUint16()
}

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
func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
    self.bootstrapMethodAttrIndex = reader.readUint16()
    self.nameAndTypeIndex = reader.readUint16()
}
