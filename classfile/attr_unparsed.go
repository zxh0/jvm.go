package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type UnparsedAttribute struct {
	name   string
	length uint32
	Info   []byte
}

func (attr *UnparsedAttribute) readInfo(reader *ClassReader) {
	attr.Info = reader.readBytes(attr.length)
}
