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
	return LineNumberTableAttribute{
		LineNumberTable: reader.readTable(func(reader *ClassReader) LineNumberTableEntry {
			return LineNumberTableEntry{
				StartPC:    reader.ReadUint16(),
				LineNumber: reader.ReadUint16(),
			}
		}).([]LineNumberTableEntry),
	}
}
