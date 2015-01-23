package classfile

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
type ConstantUtf8Info struct {
    str string
}
func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
    self.str = reader.readString()
}

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfo struct {
    stringIndex uint16
    cp          *ConstantPool
}
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
    self.stringIndex = reader.readUint16()
}
func (self *ConstantStringInfo) String() (string) {
    return self.cp.getUtf8(self.stringIndex)
}
