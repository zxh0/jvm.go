package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassInfo struct {
    nameIndex   uint16
    cp          *ConstantPool
}
func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
    self.nameIndex = reader.readUint16()
}
func (self *ConstantClassInfo) Name() (string) {
    return self.cp.getUtf8(self.nameIndex)
}

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantMemberrefInfo struct {
    classIndex          uint16
    nameAndTypeIndex    uint16
    cp                  *ConstantPool
}
func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
    self.classIndex = reader.readUint16()
    self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() (string) {
    return self.cp.getClassName(self.classIndex)
}
func (self *ConstantMemberrefInfo) Name() (string) {
    ntInfo := self.cp.getNameAndType(self.nameAndTypeIndex)
    return self.cp.getUtf8(ntInfo.nameIndex)
}
func (self *ConstantMemberrefInfo) Descriptor() (string) {
    ntInfo := self.cp.getNameAndType(self.nameAndTypeIndex)
    return self.cp.getUtf8(ntInfo.descriptorIndex)
}
func (self *ConstantMemberrefInfo) ArgCount() (uint) {
    return calcArgCount(self.Descriptor())
}

type ConstantFieldrefInfo struct {
    ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
    ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
    ConstantMemberrefInfo
}
