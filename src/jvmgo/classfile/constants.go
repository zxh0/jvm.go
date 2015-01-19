package classfile

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantIntegerInfo struct {
    val int32
}
func (self * ConstantIntegerInfo) readInfo(reader *ClassReader) {
    self.val = reader.readInt32()
}

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
