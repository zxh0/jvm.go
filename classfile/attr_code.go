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

type ExceptionTableEntry struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func readCodeAttribute(reader *ClassReader) CodeAttribute {
	return CodeAttribute{
		MaxStack:       reader.ReadUint16(),
		MaxLocals:      reader.ReadUint16(),
		Code:           reader.ReadBytes(int(reader.ReadUint32())),
		ExceptionTable: readExceptionTable(reader),
		AttributeTable: readAttributes(reader),
	}
}

func readExceptionTable(reader *ClassReader) []ExceptionTableEntry {
	return reader.readTable(func(reader *ClassReader) ExceptionTableEntry {
		return ExceptionTableEntry{
			StartPc:   reader.ReadUint16(),
			EndPc:     reader.ReadUint16(),
			HandlerPc: reader.ReadUint16(),
			CatchType: reader.ReadUint16(),
		}
	}).([]ExceptionTableEntry)
}
