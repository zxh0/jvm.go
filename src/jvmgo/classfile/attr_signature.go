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
}

func (self *SignatureAttribute) readInfo(reader *ClassReader) {
    self.signatureIndex = reader.readUint16()
}
