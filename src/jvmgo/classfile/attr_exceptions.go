package classfile

/*
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader, attrLen uint32) {
	numberOfExceptions := reader.readUint16()
	self.exceptionIndexTable = make([]uint16, numberOfExceptions)
	for i := range self.exceptionIndexTable {
		self.exceptionIndexTable[i] = reader.readUint16()
	}
}
