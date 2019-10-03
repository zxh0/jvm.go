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

func readBootstrapMethodsAttribute(reader *ClassReader) BootstrapMethodsAttribute {
	numBootstrapMethods := reader.ReadUint16()
	bootstrapMethods := make([]BootstrapMethod, numBootstrapMethods)
	for i := range bootstrapMethods {
		bootstrapMethods[i] = BootstrapMethod{
			BootstrapMethodRef: reader.ReadUint16(),
			BootstrapArguments: reader.ReadUint16s(),
		}
	}
	return BootstrapMethodsAttribute{
		BootstrapMethods: bootstrapMethods,
	}
}

type BootstrapMethod struct {
	BootstrapMethodRef uint16
	BootstrapArguments []uint16
}
