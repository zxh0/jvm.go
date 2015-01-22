package classfile

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
    maxStack        uint16
    maxLocals       uint16
    code            []byte
    exceptionTable  []*ExceptionTableEntry
    AttributeTable
}

type ExceptionTableEntry struct {
    startPc     uint16
    endPc       uint16
    handlerPc   uint16
    catchType   uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader, cp *ConstantPool) {
    self.maxStack = reader.readUint16()
    self.maxLocals = reader.readUint16()
    codeLength := reader.readUint32()
    self.code = reader.readBytes(codeLength)
    self.exceptionTable = readExceptionTable(reader)
    self.attributes = readAttributes(reader, cp)
}

func readExceptionTable(reader *ClassReader) ([]*ExceptionTableEntry) {
    exceptionTableLength := reader.readUint16()
    exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
    for i := uint16(0); i < exceptionTableLength; i++ {
        entry := &ExceptionTableEntry{}
        entry.startPc = reader.readUint16()
        entry.endPc = reader.readUint16()
        entry.handlerPc = reader.readUint16()
        entry.catchType = reader.readUint16()
        exceptionTable[i] = entry
    }
    return exceptionTable
}
