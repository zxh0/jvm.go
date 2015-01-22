package classfile

type AttributeTable struct {
    attributes []AttributeInfo
}

func (self *AttributeTable) CodeAttribute() (*CodeAttribute) {
    for _, attrInfo := range self.attributes {
        switch attrInfo.(type) {
            case *CodeAttribute: return attrInfo.(*CodeAttribute)
        }
    }
    // todo
    return nil
}
