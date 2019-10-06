package classfile

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type MemberInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributeTable
}

// read field or method table
func readMembers(reader *ClassReader) []MemberInfo {
	return reader.readTable(readMember).([]MemberInfo)
}

func readMember(reader *ClassReader) MemberInfo {
	return MemberInfo{
		AccessFlags:     reader.ReadUint16(),
		NameIndex:       reader.ReadUint16(),
		DescriptorIndex: reader.ReadUint16(),
		AttributeTable:  readAttributes(reader),
	}
}
