package classfile

/*
Signature_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 signature_index;
}
*/
type SignatureAttribute struct {
	cp             *ConstantPool
	signatureIndex uint16
}

func (attr *SignatureAttribute) readInfo(reader *ClassReader) {
	attr.signatureIndex = reader.readUint16()
}

func (attr *SignatureAttribute) Signature() string {
	return attr.cp.getUtf8(attr.signatureIndex)
}
