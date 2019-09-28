package classfile

/*
EnclosingMethod_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 class_index;
    u2 method_index;
}
*/
type EnclosingMethodAttribute struct {
	ClassIndex  uint16
	MethodIndex uint16
}

func readEnclosingMethodAttribute(reader *ClassReader) EnclosingMethodAttribute {
	return EnclosingMethodAttribute{
		ClassIndex:  reader.readUint16(),
		MethodIndex: reader.readUint16(),
	}
}
