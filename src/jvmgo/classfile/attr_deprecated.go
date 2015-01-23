package classfile

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/

type DeprecatedAttribute struct {

}

func (self *DeprecatedAttribute) readInfo(reader *ClassReader, cp *ConstantPool) {
    // read nothing
}
