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

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readVersions(reader)
	self.readConstantPool(reader)
	self.AccessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.Fields = readMembers(reader, &self.ConstantPool)
	self.Methods = readMembers(reader, &self.ConstantPool)
	self.attributes = readAttributes(reader, &self.ConstantPool)
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("Bad magic!")
	}
}

func (self *ClassFile) readVersions(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	// todo check versions
}

func (self *ClassFile) readConstantPool(reader *ClassReader) {
	self.ConstantPool = ConstantPool{cf: self}
	self.ConstantPool.read(reader)
}

func (self *ClassFile) ClassName() string {
	return self.ConstantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass != 0 {
		return self.ConstantPool.getClassName(self.superClass)
	}
	return ""
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.ConstantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
