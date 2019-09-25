package classfile

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfo struct {
	cp          *ConstantPool
	stringIndex uint16
}

func (c ConstantStringInfo) String() string {
	return c.cp.getUtf8(c.stringIndex)
}

func readConstantStringInfo(reader *ClassReader,
	cp *ConstantPool) ConstantStringInfo {

	return ConstantStringInfo{
		cp:          cp,
		stringIndex: reader.readUint16(),
	}
}
