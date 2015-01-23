package classfile

//import "fmt"

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
    readInfo(reader *ClassReader, cp *ConstantPool)
}

func readAttributes(reader *ClassReader, cp *ConstantPool) ([]AttributeInfo) {
    attributesCount := reader.readUint16()
    attributes := make([]AttributeInfo, attributesCount)
    for i := uint16(0); i < attributesCount; i++ {
        attributes[i] = readAttribute(reader, cp)
    }
    return attributes
}

func readAttribute(reader *ClassReader, cp *ConstantPool) (AttributeInfo) {
    attrNameIndex := reader.readUint16()
    attrLen := reader.readUint32()
    attrName := cp.getUtf8(attrNameIndex)
    attrInfo := newAttributeInfo(attrName, attrLen)
    attrInfo.readInfo(reader, cp)
    return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32) (AttributeInfo) {
    switch attrName {
    case "Code":                        return &CodeAttribute{}
    case "ConstantValue":               return &ConstantValueAttribute{}
    case "Deprecated":                  return &DeprecatedAttribute{}
    case "Exceptions":                  return &ExceptionsAttribute{}
    case "InnerClasses":                return &InnerClassesAttribute{}
    case "LineNumberTable":             return &LineNumberTableAttribute{}
    case "RuntimeVisibleAnnotations":   return &RuntimeVisibleAnnotationsAttribute{}
    case "Signature":                   return &SignatureAttribute{}
    case "SourceFile":                  return &SourceFileAttribute{}
    case "StackMapTable":               return &UndefinedAttribute{attrLen} // todo
    default: panic("BAD attr name:" + attrName) // todo
    }
}
