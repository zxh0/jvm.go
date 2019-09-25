package classfile

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
*/
type ConstantValueAttribute struct {
	ConstantValueIndex uint16
}

func (attr *ConstantValueAttribute) readInfo(reader *ClassReader) {
	attr.ConstantValueIndex = reader.readUint16()
}
