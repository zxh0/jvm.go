package classfile

/*
RuntimeVisibleAnnotations_attribute {
    u2         attribute_name_index;
    u4         attribute_length;
    u2         num_annotations;
    annotation annotations[num_annotations];
}
*/
type RuntimeVisibleAnnotationsAttribute struct {
	info []byte
}

func (self *RuntimeVisibleAnnotationsAttribute) readInfo(reader *ClassReader, attrLen uint32) {
	self.info = reader.readBytes(attrLen)
}

func (self *RuntimeVisibleAnnotationsAttribute) Info() []byte {
	return self.info
}
