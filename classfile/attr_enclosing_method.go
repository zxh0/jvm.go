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
	cp          *ConstantPool
	classIndex  uint16
	methodIndex uint16
}

func (attr *EnclosingMethodAttribute) readInfo(reader *ClassReader) {
	attr.classIndex = reader.readUint16()
	attr.methodIndex = reader.readUint16()
}

func (attr *EnclosingMethodAttribute) ClassName() string {
	return attr.cp.getClassName(attr.classIndex)
}

func (attr *EnclosingMethodAttribute) MethodNameAndDescriptor() (string, string) {
	if attr.methodIndex > 0 {
		return attr.cp.getNameAndType(attr.methodIndex)
	} else {
		return "", ""
	}
}
