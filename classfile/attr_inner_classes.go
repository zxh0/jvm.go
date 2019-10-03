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
	Classes []InnerClassInfo
}

type InnerClassInfo struct {
	InnerClassInfoIndex   uint16
	OuterClassInfoIndex   uint16
	InnerNameIndex        uint16
	InnerClassAccessFlags uint16
}

func readInnerClassesAttribute(reader *ClassReader) InnerClassesAttribute {
	numberOfClasses := reader.ReadUint16()
	classes := make([]InnerClassInfo, numberOfClasses)
	for i := range classes {
		classes[i] = InnerClassInfo{
			InnerClassInfoIndex:   reader.ReadUint16(),
			OuterClassInfoIndex:   reader.ReadUint16(),
			InnerNameIndex:        reader.ReadUint16(),
			InnerClassAccessFlags: reader.ReadUint16(),
		}
	}
	return InnerClassesAttribute{
		Classes: classes,
	}
}
