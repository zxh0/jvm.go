package classfile

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantLongInfo struct {
    val int64
}

func readConstantLongInfo(reader *ClassReader) (*ConstantLongInfo) {
    val := reader.readInt64()
    return &ConstantLongInfo{val}
}
