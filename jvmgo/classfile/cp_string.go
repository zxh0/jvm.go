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

func (self ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}

func readConstantStringInfo(reader *ClassReader,
	cp *ConstantPool) ConstantStringInfo {

	return ConstantStringInfo{
		cp:          cp,
		stringIndex: reader.readUint16(),
	}
}
