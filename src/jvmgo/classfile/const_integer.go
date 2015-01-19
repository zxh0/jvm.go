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
