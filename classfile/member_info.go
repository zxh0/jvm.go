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
	memberCount := reader.readUint16()
	members := make([]MemberInfo, memberCount)
	for i := range members {
		members[i] = MemberInfo{}
		members[i].read(reader)
	}
	return members
}

func (mi *MemberInfo) read(reader *ClassReader) {
	mi.AccessFlags = reader.readUint16()
	mi.NameIndex = reader.readUint16()
	mi.DescriptorIndex = reader.readUint16()
	mi.attributes = readAttributes(reader)
}
