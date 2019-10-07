package classfile

type AttributeTable []AttributeInfo

/* group 1 */

func (at AttributeTable) GetCodeAttribute() (CodeAttribute, bool) {
	for _, attrInfo := range at {
		if a, ok := attrInfo.(CodeAttribute); ok {
			return a, true
		}
	}
	return CodeAttribute{}, false
}

func (at AttributeTable) GetConstantValueIndex() uint16 {
	for _, attrInfo := range at {
		if a, ok := attrInfo.(ConstantValueAttribute); ok {
			return a.ConstantValueIndex
		}
	}
	return 0
}

func (at AttributeTable) GetExceptionIndexTable() []uint16 {
	for _, attrInfo := range at {
		if a, ok := attrInfo.(ExceptionsAttribute); ok {
			return a.ExceptionIndexTable
		}
	}
	return nil
}

func (at AttributeTable) GetBootstrapMethods() []BootstrapMethod {
	for _, attrInfo := range at {
		if a, ok := attrInfo.(BootstrapMethodsAttribute); ok {
			return a.BootstrapMethods
		}
	}
	return nil
}

/* group 2 */

func (at AttributeTable) GetEnclosingMethodAttribute() (EnclosingMethodAttribute, bool) {
	for _, attrInfo := range at {
		if a, ok := attrInfo.(EnclosingMethodAttribute); ok {
			return a, true
		}
	}
	return EnclosingMethodAttribute{}, false
}

func (at AttributeTable) GetSignatureIndex() uint16 {
	for _, attrInfo := range at {
		if a, ok := attrInfo.(SignatureAttribute); ok {
			return a.SignatureIndex
		}
	}
	return 0
}

/* group 3 */

func (at AttributeTable) GetSourceFileIndex() uint16 {
	for _, attrInfo := range at {
		if a, ok := attrInfo.(SourceFileAttribute); ok {
			return a.SourceFileIndex
		}
	}
	return 0
}

func (at AttributeTable) GetLineNumberTable() []LineNumberTableEntry {
	for _, attrInfo := range at {
		if a, ok := attrInfo.(LineNumberTableAttribute); ok {
			return a.LineNumberTable
		}
	}
	return nil
}

func (at AttributeTable) GetModuleAttribute() (ModuleAttribute, bool) {
	for _, attrInfo := range at {
		if a, ok := attrInfo.(ModuleAttribute); ok {
			return a, true
		}
	}
	return ModuleAttribute{}, false
}

/* unparsed */

func (at AttributeTable) GetRuntimeVisibleAnnotationsAttributeData() []byte {
	return at.getUnparsedAttributeData(RuntimeVisibleAnnotations)
}
func (at AttributeTable) GetRuntimeVisibleParameterAnnotationsAttributeData() []byte {
	return at.getUnparsedAttributeData(RuntimeVisibleParameterAnnotations)
}
func (at AttributeTable) GetAnnotationDefaultAttributeData() []byte {
	return at.getUnparsedAttributeData(AnnotationDefault)
}

func (at AttributeTable) getUnparsedAttributeData(name string) []byte {
	for _, attrInfo := range at {
		if a, ok := attrInfo.(UnparsedAttribute); ok && a.Name == name {
			return a.Info
		}
	}
	return nil
}
