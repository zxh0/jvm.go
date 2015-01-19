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

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantFieldrefInfo struct {
    classIndex       uint16
    nameAndTypeIndex uint16
}
func (self *ConstantFieldrefInfo) readInfo(reader *ClassReader) {
    self.classIndex = reader.readUint16()
    self.nameAndTypeIndex = reader.readUint16()
}

/*
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantMethodrefInfo struct {
    classIndex       uint16
    nameAndTypeIndex uint16
}
func (self *ConstantMethodrefInfo) readInfo(reader *ClassReader) {
    self.classIndex = reader.readUint16()
    self.nameAndTypeIndex = reader.readUint16()
}

/*
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantInterfaceMethodrefInfo struct {
    classIndex       uint16
    nameAndTypeIndex uint16
}
func (self *ConstantInterfaceMethodrefInfo) readInfo(reader *ClassReader) {
    self.classIndex = reader.readUint16()
    self.nameAndTypeIndex = reader.readUint16()
}

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
    referenceKind  uint8
    referenceIndex uint16
}
func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
    self.referenceKind = reader.readUint8()
    self.referenceIndex = reader.readUint16()
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
    bootstrapMethodAttrIndex uint16
    nameAndTypeIndex         uint16
}
func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
    self.bootstrapMethodAttrIndex = reader.readUint16()
    self.nameAndTypeIndex = reader.readUint16()
}
