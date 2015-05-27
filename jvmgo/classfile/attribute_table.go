package classfile

type AttributeTable struct {
	attributes []AttributeInfo
}

/* group 1 */

func (self *AttributeTable) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

func (self *AttributeTable) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (self *AttributeTable) ExceptionsAttribute() *ExceptionsAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ExceptionsAttribute:
			return attrInfo.(*ExceptionsAttribute)
		}
	}
	return nil
}

func (self *AttributeTable) BootstrapMethodsAttribute() *BootstrapMethodsAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *BootstrapMethodsAttribute:
			return attrInfo.(*BootstrapMethodsAttribute)
		}
	}
	return nil
}

/* group 2 */

func (self *AttributeTable) EnclosingMethodAttribute() *EnclosingMethodAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *EnclosingMethodAttribute:
			return attrInfo.(*EnclosingMethodAttribute)
		}
	}
	return nil
}

func (self *AttributeTable) SignatureAttribute() *SignatureAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *SignatureAttribute:
			return attrInfo.(*SignatureAttribute)
		}
	}
	return nil
}

/* group 3 */

func (self *AttributeTable) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}

func (self *AttributeTable) LineNumberTableAttribute() *LineNumberTableAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *LineNumberTableAttribute:
			return attrInfo.(*LineNumberTableAttribute)
		}
	}
	return nil
}

/* unparsed */

func (self *AttributeTable) RuntimeVisibleAnnotationsAttributeData() []byte {
	return self.getUnparsedAttributeData("RuntimeVisibleAnnotations")
}
func (self *AttributeTable) RuntimeVisibleParameterAnnotationsAttributeData() []byte {
	return self.getUnparsedAttributeData("RuntimeVisibleParameterAnnotationsAttribute")
}
func (self *AttributeTable) AnnotationDefaultAttributeData() []byte {
	return self.getUnparsedAttributeData("AnnotationDefault")
}

func (self *AttributeTable) getUnparsedAttributeData(name string) []byte {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *UnparsedAttribute:
			unparsedAttr := attrInfo.(*UnparsedAttribute)
			if unparsedAttr.name == name {
				return unparsedAttr.info
			}
		}
	}
	return nil
}
