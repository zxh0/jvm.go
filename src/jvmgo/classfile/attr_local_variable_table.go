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
    LocalVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
    startPc         uint16
    length          uint16
    nameIndex       uint16
    descriptorIndex uint16
    index           uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
    LocalVariableTableLength := reader.readUint16()
    self.LocalVariableTable = make([]*LocalVariableTableEntry, LocalVariableTableLength)
    for i := uint16(0); i < LocalVariableTableLength; i++ {
        entry := &LocalVariableTableEntry{}
        entry.startPc = reader.readUint16()
        entry.length = reader.readUint16()
        entry.nameIndex = reader.readUint16()
        entry.descriptorIndex = reader.readUint16()
        entry.index = reader.readUint16()
        self.LocalVariableTable[i] = entry
    }
}
