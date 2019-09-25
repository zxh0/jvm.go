package classfile

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
	//magic      uint32
	minorVersion uint16
	majorVersion uint16
	ConstantPool ConstantPool
	AccessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	Fields       []MemberInfo
	Methods      []MemberInfo
	AttributeTable
}

func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readVersions(reader)
	cf.readConstantPool(reader)
	cf.AccessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	cf.superClass = reader.readUint16()
	cf.interfaces = reader.readUint16s()
	cf.Fields = readMembers(reader, &cf.ConstantPool)
	cf.Methods = readMembers(reader, &cf.ConstantPool)
	cf.attributes = readAttributes(reader, &cf.ConstantPool)
}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("Bad magic!")
	}
}

func (cf *ClassFile) readVersions(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()

	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (cf *ClassFile) readConstantPool(reader *ClassReader) {
	cf.ConstantPool = ConstantPool{cf: cf}
	cf.ConstantPool.read(reader)
}

func (cf *ClassFile) ClassName() string {
	return cf.ConstantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string {
	if cf.superClass != 0 {
		return cf.ConstantPool.getClassName(cf.superClass)
	}
	return ""
}

func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		interfaceNames[i] = cf.ConstantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
