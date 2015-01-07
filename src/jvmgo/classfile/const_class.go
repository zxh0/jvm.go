package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassInfo struct {
    nameIndex uint16
}

func readConstantClassInfo(reader *ClassReader) (*ConstantClassInfo) {
    nameIndex := reader.readUint16()
    return &ConstantClassInfo{nameIndex}
}
