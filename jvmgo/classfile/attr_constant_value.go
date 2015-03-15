package classfile

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader, attrLen uint32) {
	self.constantValueIndex = reader.readUint16()
}
