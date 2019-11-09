package classfile

// Bytecode Behaviors for Method Handles
const (
	RefGetField         = 1 // getfield C.f:T
	RefGetStatic        = 2 // getstatic C.f:T
	RefPutField         = 3 // putfield C.f:T
	RefPutStatic        = 4 // putstatic C.f:T
	RefInvokeVirtual    = 5 // invokevirtual C.m:(A*)T
	RefInvokeStatic     = 6 // invokestatic C.m:(A*)T
	RefInvokeSpecial    = 7 // invokespecial C.m:(A*)T
	RefNewInvokeSpecial = 8 // new C; dup; invokespecial C.<init>:(A*)void
	RefInvokeInterface  = 9 // invokeinterface C.m:(A*)T
)

/*
CONSTANT_Dynamic_info {
	u1 tag;
	u2 bootstrap_method_attr_index;
	u2 name_and_type_index;
}
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
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

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfo struct {
	DescriptorIndex uint16
}

func readConstantInvokeDynamicInfo(reader *ClassReader) ConstantInvokeDynamicInfo {
	return ConstantInvokeDynamicInfo{
		BootstrapMethodAttrIndex: reader.ReadUint16(),
		NameAndTypeIndex:         reader.ReadUint16(),
	}
}

func readConstantMethodHandleInfo(reader *ClassReader) ConstantMethodHandleInfo {
	return ConstantMethodHandleInfo{
		ReferenceKind:  reader.ReadUint8(),
		ReferenceIndex: reader.ReadUint16(),
	}
}

func readConstantMethodTypeInfo(reader *ClassReader) ConstantMethodTypeInfo {
	return ConstantMethodTypeInfo{
		DescriptorIndex: reader.ReadUint16(),
	}
}
