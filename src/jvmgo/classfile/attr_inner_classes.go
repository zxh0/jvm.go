package classfile

/*
InnerClasses_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_classes;
    {   u2 inner_class_info_index;
        u2 outer_class_info_index;
        u2 inner_name_index;
        u2 inner_class_access_flags;
    } classes[number_of_classes];
}
*/
type InnerClassesAttribute struct {
    classes []*InnerClassInfo
}
func (self *InnerClassesAttribute) readInfo(reader *ClassReader, cp *ConstantPool) {
    numberOfClasses := reader.readUint16()
    self.classes = make([]*InnerClassInfo, numberOfClasses)
    for i := range self.classes {
        self.classes[i] = readInnerClassInfo(reader)
    }
}

type InnerClassInfo struct {
    innerClassInfoIndex     uint16
    outerClassInfoIndex     uint16
    innerNameIndex          uint16
    innerClassAccessFlags   uint16
}
func readInnerClassInfo(reader *ClassReader) (*InnerClassInfo) {
    info := &InnerClassInfo{}
    info.innerClassInfoIndex = reader.readUint16()
    info.outerClassInfoIndex = reader.readUint16()
    info.innerNameIndex = reader.readUint16()
    info.innerClassAccessFlags = reader.readUint16()
    return info
}
