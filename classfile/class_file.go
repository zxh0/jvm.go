package classfile

import (
	"fmt"
)

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
	MinorVersion uint16
	MajorVersion uint16
	ConstantPool []ConstantInfo
	AccessFlags  uint16
	ThisClass    uint16
	SuperClass   uint16
	Interfaces   []uint16
	Fields       []MemberInfo
	Methods      []MemberInfo
	AttributeTable
}

func (cf *ClassFile) read(reader *ClassReader) {
	reader.cf = cf
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersions(reader)
	cf.ConstantPool = readConstantPool(reader)
	cf.AccessFlags = reader.ReadUint16()
	cf.ThisClass = reader.ReadUint16()
	cf.SuperClass = reader.ReadUint16()
	cf.Interfaces = reader.readUint16s()
	cf.Fields = readMembers(reader)
	cf.Methods = readMembers(reader)
	cf.attributes = readAttributes(reader)
}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.ReadUint32()
	if magic != 0xCAFEBABE {
		panic("Bad magic!") // TODO
	}
}

func (cf *ClassFile) readAndCheckVersions(reader *ClassReader) {
	cf.MinorVersion = reader.ReadUint16()
	cf.MajorVersion = reader.ReadUint16()

	switch cf.MajorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52,
		53, 54, 55, 56, 57:
		if cf.MinorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (cf *ClassFile) GetThisClassName() string {
	return cf.GetClassNameOf(cf.ThisClass)
}
func (cf *ClassFile) GetSuperClassName() string {
	return cf.GetClassNameOf(cf.SuperClass)
}
func (cf *ClassFile) GetInterfaceNames() []string {
	interfaceNames := make([]string, len(cf.Interfaces))
	for i, cpIndex := range cf.Interfaces {
		interfaceNames[i] = cf.GetClassNameOf(cpIndex)
	}
	return interfaceNames
}

func (cf *ClassFile) GetUTF8(index uint16) string {
	return cf.getUtf8(index)
}
func (cf *ClassFile) GetConstantInfo(index uint16) ConstantInfo {
	return cf.getConstantInfo(index)
}

func (cf *ClassFile) GetClassNameOf(cpIndex uint16) string {
	if cpIndex == 0 {
		return ""
	}
	classInfo := cf.getConstantInfo(cpIndex).(ConstantClassInfo)
	return cf.getUtf8(classInfo.NameIndex)
}

func (cf *ClassFile) getUtf8(index uint16) string {
	if index == 0 {
		return ""
	}
	return cf.getConstantInfo(index).(string)
}

func (cf *ClassFile) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := cf.ConstantPool[index]; cpInfo == nil {
		panic(fmt.Errorf("invalid constant pool index: %d", index))
	} else {
		return cpInfo
	}
}
