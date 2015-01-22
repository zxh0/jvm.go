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
    numberOfExceptions  uint16
    exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader, cp *ConstantPool) {
    self.numberOfExceptions = reader.readUint16()
    // todo
}
