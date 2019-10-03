package classfile

/*
Signature_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 signature_index;
}
*/
type SignatureAttribute struct {
	SignatureIndex uint16
}

func readSignatureAttribute(reader *ClassReader) SignatureAttribute {
	return SignatureAttribute{
		SignatureIndex: reader.ReadUint16(),
	}
}
