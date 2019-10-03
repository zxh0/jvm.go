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

func readConstantValueAttribute(reader *ClassReader) ConstantValueAttribute {
	return ConstantValueAttribute{
		ConstantValueIndex: reader.ReadUint16(),
	}
}
