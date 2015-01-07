package classfile

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
type ConstantUtf8Info struct {
    str string
}

func readConstantUtf8Info(reader *ClassReader) (*ConstantUtf8Info) {
    str := reader.readString()
    return &ConstantUtf8Info{str}
}
