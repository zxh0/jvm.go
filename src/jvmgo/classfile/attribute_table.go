package classfile

type AttributeTable struct {
	attributes []AttributeInfo
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

func (self *AttributeTable) EnclosingMethodAttribute() *EnclosingMethodAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *EnclosingMethodAttribute:
			return attrInfo.(*EnclosingMethodAttribute)
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

func (self *AttributeTable) LineNumberTableAttribute() *LineNumberTableAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *LineNumberTableAttribute:
			return attrInfo.(*LineNumberTableAttribute)
		}
	}
	return nil
}

func (self *AttributeTable) RuntimeVisibleAnnotationsAttribute() *RuntimeVisibleAnnotationsAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *RuntimeVisibleAnnotationsAttribute:
			return attrInfo.(*RuntimeVisibleAnnotationsAttribute)
		}
	}
	return nil
}

func (self *AttributeTable) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}
