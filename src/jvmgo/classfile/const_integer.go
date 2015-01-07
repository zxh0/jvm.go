package classfile

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantIntegerInfo struct {
    val int32
}

func readConstantIntegerInfo(reader *ClassReader) (*ConstantIntegerInfo) {
    val := reader.readInt32()
    return &ConstantIntegerInfo{val}
}
