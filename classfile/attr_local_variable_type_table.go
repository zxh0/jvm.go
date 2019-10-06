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
	return LocalVariableTypeTableAttribute{
		LocalVariableTypeTable: reader.readTable(func(reader *ClassReader) LocalVariableTypeTableEntry {
			return LocalVariableTypeTableEntry{
				StartPc:        reader.ReadUint16(),
				Length:         reader.ReadUint16(),
				NameIndex:      reader.ReadUint16(),
				SignatureIndex: reader.ReadUint16(),
				Index:          reader.ReadUint16(),
			}
		}).([]LocalVariableTypeTableEntry),
	}
}
