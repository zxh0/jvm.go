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

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
    self.sourceFileIndex = reader.readUint16()
}
