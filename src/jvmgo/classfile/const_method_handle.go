package classfile

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
    referenceKind  uint8
    referenceIndex uint16
}

func readConstantMethodHandleInfo(reader *ClassReader) (*ConstantMethodHandleInfo) {
    referenceKind := reader.readUint8()
    referenceIndex := reader.readUint16()
    return &ConstantMethodHandleInfo{referenceKind, referenceIndex}
}
