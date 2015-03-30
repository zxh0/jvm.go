package classfile

/*
Signature_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 signature_index;
}
*/
type SignatureAttribute struct {
	signatureIndex uint16
	cp             *ConstantPool
}

func (self *SignatureAttribute) readInfo(reader *ClassReader, attrLen uint32) {
	self.signatureIndex = reader.readUint16()
}

func (self *SignatureAttribute) SignatureName() string {
	return self.cp.getUtf8(self.signatureIndex)
}
