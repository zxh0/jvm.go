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
    classIndex  uint16
    methodIndex uint16
}

func (self *EnclosingMethodAttribute) readInfo(reader *ClassReader) {
    self.classIndex = reader.readUint16()
    self.methodIndex = reader.readUint16()
}
