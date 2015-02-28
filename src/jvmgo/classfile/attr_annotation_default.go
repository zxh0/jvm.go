package classfile

/*
AnnotationDefault_attribute {
    u2            attribute_name_index;
    u4            attribute_length;
    element_value default_value;
}
*/
type AnnotationDefaultAttribute struct {
	elementValue *ElementValue
}

func (self *AnnotationDefaultAttribute) readInfo(reader *ClassReader, attrLen uint32) {
	self.elementValue = readElementValue(reader)
}
