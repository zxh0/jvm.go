package classfile

/*
LocalVariableTypeTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_type_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 signature_index;
        u2 index;
    } local_variable_type_table[local_variable_type_table_length];
}
*/
type LocalVariableTypeTableAttribute struct {
	LocalVariableTypeTable []LocalVariableTypeTableEntry
}

type LocalVariableTypeTableEntry struct {
	StartPc        uint16
	Length         uint16
	NameIndex      uint16
	SignatureIndex uint16
	Index          uint16
}

func readLocalVariableTypeTableAttribute(reader *ClassReader) LocalVariableTypeTableAttribute {
	tableLength := reader.readUint16()
	localVariableTypeTable := make([]LocalVariableTypeTableEntry, tableLength)
	for i := range localVariableTypeTable {
		localVariableTypeTable[i] = LocalVariableTypeTableEntry{
			StartPc:        reader.readUint16(),
			Length:         reader.readUint16(),
			NameIndex:      reader.readUint16(),
			SignatureIndex: reader.readUint16(),
			Index:          reader.readUint16(),
		}
	}
	return LocalVariableTypeTableAttribute{
		LocalVariableTypeTable: localVariableTypeTable,
	}
}
