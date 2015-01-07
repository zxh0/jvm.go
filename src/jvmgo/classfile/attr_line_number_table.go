package classfile

/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number; 
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableAttribute struct {
    lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
    startPc     uint16
    lineNumber  uint16
}

func readLineNumberTableAttribute(reader *ClassReader) (*LineNumberTableAttribute) {
    lineNumberTableLength := reader.readUint16()
    lineNumberTable := make([]*LineNumberTableEntry, lineNumberTableLength)
    for i := uint16(0); i < lineNumberTableLength; i++ {
        entry := &LineNumberTableEntry{}
        entry.startPc = reader.readUint16()
        entry.lineNumber = reader.readUint16()
        lineNumberTable[i] = entry
    }

    return &LineNumberTableAttribute{lineNumberTable}
}
