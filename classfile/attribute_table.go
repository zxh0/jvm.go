package classfile

type AttributeTable []AttributeInfo

/* group 1 */

func (at AttributeTable) GetCodeAttribute() *CodeAttribute {
	for _, attrInfo := range at {
		switch attrInfo.(type) {
		case CodeAttribute:
			ca := attrInfo.(CodeAttribute)
			return &ca
		}
	}
	return nil
}

func (at AttributeTable) GetConstantValueIndex() uint16 {
	for _, attrInfo := range at {
		switch attrInfo.(type) {
		case ConstantValueAttribute:
			return attrInfo.(ConstantValueAttribute).ConstantValueIndex
		}
	}
	return 0
}

func (at AttributeTable) GetExceptionIndexTable() []uint16 {
	for _, attrInfo := range at {
		switch attrInfo.(type) {
		case ExceptionsAttribute:
			return attrInfo.(ExceptionsAttribute).ExceptionIndexTable
		}
	}
	return nil
}

func (at AttributeTable) GetBootstrapMethods() []BootstrapMethod {
	for _, attrInfo := range at {
		switch attrInfo.(type) {
		case BootstrapMethodsAttribute:
			return attrInfo.(BootstrapMethodsAttribute).BootstrapMethods
		}
	}
	return nil
}

/* group 2 */

func (at AttributeTable) GetEnclosingMethodAttribute() (EnclosingMethodAttribute, bool) {
	for _, attrInfo := range at {
		switch attrInfo.(type) {
		case EnclosingMethodAttribute:
			return attrInfo.(EnclosingMethodAttribute), true
		}
	}
	return EnclosingMethodAttribute{}, false
}

func (at AttributeTable) GetSignatureIndex() uint16 {
	for _, attrInfo := range at {
		switch attrInfo.(type) {
		case SignatureAttribute:
			return attrInfo.(SignatureAttribute).SignatureIndex
		}
	}
	return 0
}

/* group 3 */

func (at AttributeTable) GetSourceFileIndex() uint16 {
	for _, attrInfo := range at {
		switch attrInfo.(type) {
		case SourceFileAttribute:
			return attrInfo.(SourceFileAttribute).SourceFileIndex
		}
	}
	return 0
}

func (at AttributeTable) GetLineNumberTable() []LineNumberTableEntry {
	for _, attrInfo := range at {
		switch attrInfo.(type) {
		case LineNumberTableAttribute:
			return attrInfo.(LineNumberTableAttribute).LineNumberTable
		}
	}
	return nil
}

/* unparsed */

func (at AttributeTable) GetRuntimeVisibleAnnotationsAttributeData() []byte {
	return at.getUnparsedAttributeData("RuntimeVisibleAnnotations")
}
func (at AttributeTable) GetRuntimeVisibleParameterAnnotationsAttributeData() []byte {
	return at.getUnparsedAttributeData("RuntimeVisibleParameterAnnotationsAttribute")
}
func (at AttributeTable) GetAnnotationDefaultAttributeData() []byte {
	return at.getUnparsedAttributeData("AnnotationDefault")
}

func (at AttributeTable) getUnparsedAttributeData(name string) []byte {
	for _, attrInfo := range at {
		switch attrInfo.(type) {
		case UnparsedAttribute:
			unparsedAttr := attrInfo.(UnparsedAttribute)
			if unparsedAttr.Name == name {
				return unparsedAttr.Info
			}
		}
	}
	return nil
}
