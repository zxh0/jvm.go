package classfile

import (
	"fmt"

	"github.com/zxh0/jvm.go/vmutils"
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
	cf.AttributeTable = readAttributes(reader)
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
	return cf.GetClassName(cf.ThisClass)
}
func (cf *ClassFile) GetSuperClassName() string {
	return cf.GetClassName(cf.SuperClass)
}
func (cf *ClassFile) GetInterfaceNames() []string {
	return cf.GetClassNames(cf.Interfaces)
}

func (cf *ClassFile) GetNameAndType(cpIndex uint16) (name, _type string) {
	if cpIndex > 0 {
		ntInfo := cf.getConstantInfo(cpIndex).(ConstantNameAndTypeInfo)
		name = cf.GetUTF8(ntInfo.NameIndex)
		_type = cf.GetUTF8(ntInfo.DescriptorIndex)
	}
	return
}

func (cf *ClassFile) GetClassName(cpIndex uint16) string {
	if cpIndex == 0 {
		return ""
	}
	classInfo := cf.getConstantInfo(cpIndex).(ConstantClassInfo)
	return cf.GetUTF8(classInfo.NameIndex)
}
func (cf *ClassFile) GetPackageName(cpIndex uint16) string {
	if cpIndex == 0 {
		return ""
	}
	pkgInfo := cf.getConstantInfo(cpIndex).(ConstantPackageInfo)
	return cf.GetUTF8(pkgInfo.NameIndex)
}
func (cf *ClassFile) GetModuleName(cpIndex uint16) string {
	if cpIndex == 0 {
		return ""
	}
	modInfo := cf.getConstantInfo(cpIndex).(ConstantModuleInfo)
	return cf.GetUTF8(modInfo.NameIndex)
}

func (cf *ClassFile) GetClassNames(cpIndexes []uint16) []string {
	ss := make([]string, len(cpIndexes))
	for i, cpIndex := range cpIndexes {
		ss[i] = cf.GetClassName(cpIndex)
	}
	return ss
}
func (cf *ClassFile) GetModuleNames(cpIndexes []uint16) []string {
	ss := make([]string, len(cpIndexes))
	for i, cpIndex := range cpIndexes {
		ss[i] = cf.GetModuleName(cpIndex)
	}
	return ss
}

func (cf *ClassFile) GetRawUTF8(cpIndex uint16) string {
	if cpIndex == 0 {
		return ""
	}
	rawBytes := cf.getConstantInfo(cpIndex).([]byte)
	return string(rawBytes)
}
func (cf *ClassFile) GetUTF8(cpIndex uint16) string {
	if cpIndex == 0 {
		return ""
	}
	bytes := cf.getConstantInfo(cpIndex).([]byte)
	return vmutils.DecodeMUTF8(bytes)
}

func (cf *ClassFile) getConstantInfo(cpIndex uint16) ConstantInfo {
	if cpInfo := cf.ConstantPool[cpIndex]; cpInfo == nil {
		panic(fmt.Errorf("invalid constant pool index: %d", cpIndex))
	} else {
		return cpInfo
	}
}
