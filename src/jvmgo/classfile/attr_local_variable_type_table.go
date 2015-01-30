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
    localVariableTypeTable []*LocalVariableTypeTableEntry
}

type LocalVariableTypeTableEntry struct {
    startPc         uint16
    length          uint16
    nameIndex       uint16
    signatureIndex  uint16
    index           uint16
}

func (self *LocalVariableTypeTableAttribute) readInfo(reader *ClassReader, cp *ConstantPool) {
    localVariableTypeTableLength := reader.readUint16()
    self.localVariableTypeTable = make([]*LocalVariableTypeTableEntry, localVariableTypeTableLength)
    for i := uint16(0); i < localVariableTypeTableLength; i++ {
        entry := &LocalVariableTypeTableEntry{}
        entry.startPc = reader.readUint16()
        entry.length = reader.readUint16()
        entry.nameIndex = reader.readUint16()
        entry.signatureIndex = reader.readUint16()
        entry.index = reader.readUint16()
        self.localVariableTypeTable[i] = entry
    }
}
