package classfile

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantDoubleInfo struct {
    val float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
    self.val = reader.readFloat64()
}
