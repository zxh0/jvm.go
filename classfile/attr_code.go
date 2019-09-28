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
	MaxStack       uint16
	MaxLocals      uint16
	Code           []byte
	ExceptionTable []ExceptionTableEntry
	AttributeTable
}

func readCodeAttribute(reader *ClassReader) CodeAttribute {
	return CodeAttribute{
		MaxStack:       reader.readUint16(),
		MaxLocals:      reader.readUint16(),
		Code:           reader.readBytes(reader.readUint32()),
		ExceptionTable: readExceptionTable(reader),
		AttributeTable: AttributeTable{
			readAttributes(reader),
		},
	}
}

type ExceptionTableEntry struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func readExceptionTable(reader *ClassReader) []ExceptionTableEntry {
	tableLength := reader.readUint16()
	exceptionTable := make([]ExceptionTableEntry, tableLength)
	for i := range exceptionTable {
		exceptionTable[i] = ExceptionTableEntry{
			StartPc:   reader.readUint16(),
			EndPc:     reader.readUint16(),
			HandlerPc: reader.readUint16(),
			CatchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
