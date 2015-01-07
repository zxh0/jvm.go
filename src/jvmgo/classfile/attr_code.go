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
    attributes      []AttributeInfo
}

type ExceptionTableEntry struct {
    startPc     uint16
    endPc       uint16
    handlerPc   uint16
    catchType   uint16
}

func readCodeAttribute(reader *ClassReader, cp *ConstantPool) (*CodeAttribute) {
    codeAttr := &CodeAttribute{}
    codeAttr.maxStack = reader.readUint16()
    codeAttr.maxLocals = reader.readUint16()
    codeLength := reader.readUint32()
    codeAttr.code = reader.readBytes(codeLength)
    codeAttr.exceptionTable = readExceptionTable(reader)
    codeAttr.attributes = readAttributes(reader, cp)
    return codeAttr
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
