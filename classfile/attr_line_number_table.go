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
	LineNumberTable []LineNumberTableEntry
}

type LineNumberTableEntry struct {
	StartPC    uint16
	LineNumber uint16
}

func readLineNumberTableAttribute(reader *ClassReader) LineNumberTableAttribute {
	tableLength := reader.readUint16()
	lineNumberTable := make([]LineNumberTableEntry, tableLength)
	for i := range lineNumberTable {
		lineNumberTable[i] = LineNumberTableEntry{
			StartPC:    reader.readUint16(),
			LineNumber: reader.readUint16(),
		}
	}
	return LineNumberTableAttribute{
		LineNumberTable: lineNumberTable,
	}
}
