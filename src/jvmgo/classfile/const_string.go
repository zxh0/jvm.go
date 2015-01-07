package classfile

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfo struct {
    stringIndex uint16
}

func readConstantStringInfo(reader *ClassReader) (*ConstantStringInfo) {
    stringIndex := reader.readUint16()
    return &ConstantStringInfo{stringIndex}
}
