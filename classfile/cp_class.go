package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassInfo struct {
	NameIndex uint16
}

func readConstantClassInfo(reader *ClassReader) ConstantClassInfo {
	return ConstantClassInfo{
		NameIndex: reader.readUint16(),
	}
}
