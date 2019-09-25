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
	cp              *ConstantPool
	AccessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	AttributeTable
}

// read field or method table
func readMembers(reader *ClassReader, cp *ConstantPool) []MemberInfo {
	memberCount := reader.readUint16()
	members := make([]MemberInfo, memberCount)
	for i := range members {
		members[i] = MemberInfo{cp: cp}
		members[i].read(reader)
	}
	return members
}

func (mi *MemberInfo) read(reader *ClassReader) {
	mi.AccessFlags = reader.readUint16()
	mi.nameIndex = reader.readUint16()
	mi.descriptorIndex = reader.readUint16()
	mi.attributes = readAttributes(reader, mi.cp)
}

func (mi *MemberInfo) Name() string {
	return mi.cp.getUtf8(mi.nameIndex)
}
func (mi *MemberInfo) Descriptor() string {
	return mi.cp.getUtf8(mi.descriptorIndex)
}
func (mi *MemberInfo) Signature() string {
	signatureAttr := mi.SignatureAttribute()
	if signatureAttr != nil {
		return signatureAttr.Signature()
	}
	return ""
}
