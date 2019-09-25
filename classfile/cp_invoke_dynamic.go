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

func (self ConstantInvokeDynamicInfo) NameAndType() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

// todo
func (self ConstantInvokeDynamicInfo) BootstrapMethodInfo() (uint16, []uint16) {
	bmAttr := self.cp.cf.BootstrapMethodsAttribute()
	bm := bmAttr.bootstrapMethods[self.bootstrapMethodAttrIndex]

	return bm.bootstrapMethodRef, bm.bootstrapArguments
}

func readConstantInvokeDynamicInfo(reader *ClassReader,
	cp *ConstantPool) ConstantInvokeDynamicInfo {

	return ConstantInvokeDynamicInfo{
		cp:                       cp,
		bootstrapMethodAttrIndex: reader.readUint16(),
		nameAndTypeIndex:         reader.readUint16(),
	}
}

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func readConstantMethodHandleInfo(reader *ClassReader) ConstantMethodHandleInfo {
	return ConstantMethodHandleInfo{
		ReferenceKind:  reader.readUint8(),
		ReferenceIndex: reader.readUint16(),
	}
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

func readConstantMethodTypeInfo(reader *ClassReader) ConstantMethodTypeInfo {
	return ConstantMethodTypeInfo{
		descriptorIndex: reader.readUint16(),
	}
}
