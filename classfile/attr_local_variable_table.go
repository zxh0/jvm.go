package classfile

/*
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
*/
type LocalVariableTableAttribute struct {
	LocalVariableTable []LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	StartPc         uint16
	Length          uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Index           uint16
}

func readLocalVariableTableAttribute(reader *ClassReader) LocalVariableTableAttribute {
	tableLength := reader.readUint16()
	localVariableTable := make([]LocalVariableTableEntry, tableLength)
	for i := range localVariableTable {
		localVariableTable[i] = LocalVariableTableEntry{
			StartPc:         reader.readUint16(),
			Length:          reader.readUint16(),
			NameIndex:       reader.readUint16(),
			DescriptorIndex: reader.readUint16(),
			Index:           reader.readUint16(),
		}
	}
	return LocalVariableTableAttribute{
		LocalVariableTable: localVariableTable,
	}
}
