package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type UndefinedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (self *UndefinedAttribute) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.length)
}

func (self *UndefinedAttribute) Info() []byte {
	return self.info
}
