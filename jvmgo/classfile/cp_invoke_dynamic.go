package classfile

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
	cp                       *ConstantPool
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootstrapMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantInvokeDynamicInfo) NameAndType() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

// todo
func (self *ConstantInvokeDynamicInfo) BootstrapMethodInfo() (uint16, []uint16) {
	bmAttr := self.cp.cf.BootstrapMethodsAttribute()
	bm := bmAttr.bootstrapMethods[self.bootstrapMethodAttrIndex]

	return bm.bootstrapMethodRef, bm.bootstrapArguments
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
func (self *ConstantMethodHandleInfo) ReferenceKind() uint8 {
	return self.referenceKind
}
func (self *ConstantMethodHandleInfo) ReferenceIndex() uint16 {
	return self.referenceIndex
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
