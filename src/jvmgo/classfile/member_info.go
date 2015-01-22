package classfile

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type FieldInfo struct {
    accessFlags     uint16
    nameIndex       uint16
    descriptorIndex uint16
    attributes      []AttributeInfo
}

/*
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type MethodInfo struct {
    FieldInfo
}

func (self *FieldInfo) read(reader *ClassReader, cp *ConstantPool) {
    self.accessFlags = reader.readUint16()
    self.nameIndex = reader.readUint16()
    self.descriptorIndex = reader.readUint16()
    self.attributes = readAttributes(reader, cp)
}
