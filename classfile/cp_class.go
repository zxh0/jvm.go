package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassInfo struct {
	cp        *ConstantPool
	nameIndex uint16
}

func (c ConstantClassInfo) Name() string {
	return c.cp.getUtf8(c.nameIndex)
}

func readConstantClassInfo(reader *ClassReader,
	cp *ConstantPool) ConstantClassInfo {

	return ConstantClassInfo{
		cp:        cp,
		nameIndex: reader.readUint16(),
	}
}
