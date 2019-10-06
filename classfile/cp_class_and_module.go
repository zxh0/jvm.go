package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}

CONSTANT_Module_info {
    u1 tag;
    u2 name_index;
}

CONSTANT_Package_info {
    u1 tag;
    u2 name_index;
}
*/

type ConstantClassInfo constantWithNameIdx
type ConstantModuleInfo constantWithNameIdx
type ConstantPackageInfo constantWithNameIdx

type constantWithNameIdx struct {
	NameIndex uint16
}

func readConstantClassInfo(reader *ClassReader) ConstantClassInfo {
	return ConstantClassInfo(readConstantWithNameIdx(reader))
}
func readConstantModuleInfo(reader *ClassReader) ConstantModuleInfo {
	return ConstantModuleInfo(readConstantWithNameIdx(reader))
}
func readConstantPackageInfo(reader *ClassReader) ConstantPackageInfo {
	return ConstantPackageInfo(readConstantWithNameIdx(reader))
}

func readConstantWithNameIdx(reader *ClassReader) constantWithNameIdx {
	return constantWithNameIdx{
		NameIndex: reader.ReadUint16(),
	}
}
