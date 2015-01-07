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

func readFieldInfo(reader *ClassReader, cp *ConstantPool) (*FieldInfo) {
    fieldInfo := &FieldInfo{}
    fieldInfo.accessFlags = reader.readUint16()
    fieldInfo.nameIndex = reader.readUint16()
    fieldInfo.descriptorIndex = reader.readUint16()
    fieldInfo.attributes = readAttributes(reader, cp)
    return fieldInfo
}

func readMethodInfo(reader *ClassReader, cp *ConstantPool) (*MethodInfo) {
    methodInfo := &MethodInfo{}
    methodInfo.accessFlags = reader.readUint16()
    methodInfo.nameIndex = reader.readUint16()
    methodInfo.descriptorIndex = reader.readUint16()
    methodInfo.attributes = readAttributes(reader, cp)
    return methodInfo
}
