package classfile

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantFloatInfo struct {
    val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
    self.val = reader.readFloat32()
}
