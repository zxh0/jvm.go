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

func (self *SignatureAttribute) readInfo(reader *ClassReader) {
	self.signatureIndex = reader.readUint16()
}

func (self *SignatureAttribute) Signature() string {
	return self.cp.getUtf8(self.signatureIndex)
}
