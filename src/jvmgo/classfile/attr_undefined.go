package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type UndefinedAttribute struct {
	
}

func (self *UndefinedAttribute) readInfo(reader *ClassReader, attrLen uint32) {
	for i := uint32(0); i < attrLen; i++ {
		reader.readUint8()
	}
}
