package classfile

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/
type SourceFileAttribute struct {
    sourceFileIndex uint16
}

func readSourceFileAttribute(reader *ClassReader) (*SourceFileAttribute) {
    sourceFileIndex := reader.readUint16()
    return &SourceFileAttribute{sourceFileIndex}
}
