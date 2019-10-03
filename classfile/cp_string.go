package classfile

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfo struct {
	StringIndex uint16
}

func readConstantStringInfo(reader *ClassReader) ConstantStringInfo {
	return ConstantStringInfo{
		StringIndex: reader.ReadUint16(),
	}
}
