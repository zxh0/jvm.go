package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type UndefinedAttribute struct {
    attributeLength uint32
}

func (self *UndefinedAttribute) readInfo(reader *ClassReader) {
    for i := uint32(0); i < self.attributeLength; i++ {
        reader.readUint8()
    }
}
