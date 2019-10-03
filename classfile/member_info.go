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
	memberCount := reader.ReadUint16()
	members := make([]MemberInfo, memberCount)
	for i := range members {
		members[i] = MemberInfo{}
		members[i].read(reader)
	}
	return members
}

func (mi *MemberInfo) read(reader *ClassReader) {
	mi.AccessFlags = reader.ReadUint16()
	mi.NameIndex = reader.ReadUint16()
	mi.DescriptorIndex = reader.ReadUint16()
	mi.attributes = readAttributes(reader)
}
