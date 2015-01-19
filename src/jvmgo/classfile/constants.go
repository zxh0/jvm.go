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
}
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
    self.stringIndex = reader.readUint16()
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

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfo struct {
    descriptorIndex uint16
}
func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
    self.descriptorIndex = reader.readUint16()
}

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/
type ConstantNameAndTypeInfo struct {
    nameIndex       uint16
    descriptorIndex uint16
}
func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
    self.nameIndex = reader.readUint16()
    self.descriptorIndex = reader.readUint16()
}
