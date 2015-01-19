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
