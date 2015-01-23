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
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
    self.val = reader.readInt32()
}
func (self *ConstantIntegerInfo) Value() (int32) {
    return self.val
}

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
func (self *ConstantFloatInfo) Value() (float32) {
    return self.val
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantLongInfo struct {
    val int64
}
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
    self.val = reader.readInt64()
}
func (self *ConstantLongInfo) Value() (int64) {
    return self.val
}

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
func (self *ConstantDoubleInfo) Value() (float64) {
    return self.val
}
