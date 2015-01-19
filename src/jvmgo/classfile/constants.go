package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassInfo struct {
    nameIndex uint16
}
func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
    self.nameIndex = reader.readUint16()
}
