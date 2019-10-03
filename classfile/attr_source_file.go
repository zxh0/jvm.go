package classfile

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/
type SourceFileAttribute struct {
	SourceFileIndex uint16
}

func readSourceFileAttribute(reader *ClassReader) SourceFileAttribute {
	return SourceFileAttribute{SourceFileIndex: reader.ReadUint16()}
}
