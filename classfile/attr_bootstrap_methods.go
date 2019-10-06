package classfile

/*
BootstrapMethods_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 num_bootstrap_methods;
    {   u2 bootstrap_method_ref;
        u2 num_bootstrap_arguments;
        u2 bootstrap_arguments[num_bootstrap_arguments];
    } bootstrap_methods[num_bootstrap_methods];
}
*/
type BootstrapMethodsAttribute struct {
	BootstrapMethods []BootstrapMethod
}

type BootstrapMethod struct {
	BootstrapMethodRef uint16
	BootstrapArguments []uint16
}

func readBootstrapMethodsAttribute(reader *ClassReader) BootstrapMethodsAttribute {
	return BootstrapMethodsAttribute{
		BootstrapMethods: reader.readTable(func(reader *ClassReader) BootstrapMethod {
			return BootstrapMethod{
				BootstrapMethodRef: reader.ReadUint16(),
				BootstrapArguments: reader.readUint16s(),
			}
		}).([]BootstrapMethod),
	}
}
