package classfile

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantFloatInfo struct {
    val float32
}

func readConstantFloatInfo(reader *ClassReader) (*ConstantFloatInfo) {
    val := reader.readFloat32()
    return &ConstantFloatInfo{val}
}
