package classfile

type AttributeTable struct {
	attributes []AttributeInfo
}

/* group 1 */

func (at *AttributeTable) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range at.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

func (at *AttributeTable) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range at.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (at *AttributeTable) ExceptionsAttribute() *ExceptionsAttribute {
	for _, attrInfo := range at.attributes {
		switch attrInfo.(type) {
		case *ExceptionsAttribute:
			return attrInfo.(*ExceptionsAttribute)
		}
	}
	return nil
}

func (at *AttributeTable) BootstrapMethodsAttribute() *BootstrapMethodsAttribute {
	for _, attrInfo := range at.attributes {
		switch attrInfo.(type) {
		case *BootstrapMethodsAttribute:
			return attrInfo.(*BootstrapMethodsAttribute)
		}
	}
	return nil
}

/* group 2 */

func (at *AttributeTable) EnclosingMethodAttribute() *EnclosingMethodAttribute {
	for _, attrInfo := range at.attributes {
		switch attrInfo.(type) {
		case *EnclosingMethodAttribute:
			return attrInfo.(*EnclosingMethodAttribute)
		}
	}
	return nil
}

func (at *AttributeTable) SignatureAttribute() *SignatureAttribute {
	for _, attrInfo := range at.attributes {
		switch attrInfo.(type) {
		case *SignatureAttribute:
			return attrInfo.(*SignatureAttribute)
		}
	}
	return nil
}

/* group 3 */

func (at *AttributeTable) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range at.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}

func (at *AttributeTable) LineNumberTableAttribute() *LineNumberTableAttribute {
	for _, attrInfo := range at.attributes {
		switch attrInfo.(type) {
		case *LineNumberTableAttribute:
			return attrInfo.(*LineNumberTableAttribute)
		}
	}
	return nil
}

/* unparsed */

func (at *AttributeTable) RuntimeVisibleAnnotationsAttributeData() []byte {
	return at.getUnparsedAttributeData("RuntimeVisibleAnnotations")
}
func (at *AttributeTable) RuntimeVisibleParameterAnnotationsAttributeData() []byte {
	return at.getUnparsedAttributeData("RuntimeVisibleParameterAnnotationsAttribute")
}
func (at *AttributeTable) AnnotationDefaultAttributeData() []byte {
	return at.getUnparsedAttributeData("AnnotationDefault")
}

func (at *AttributeTable) getUnparsedAttributeData(name string) []byte {
	for _, attrInfo := range at.attributes {
		switch attrInfo.(type) {
		case *UnparsedAttribute:
			unparsedAttr := attrInfo.(*UnparsedAttribute)
			if unparsedAttr.name == name {
				return unparsedAttr.Info
			}
		}
	}
	return nil
}
