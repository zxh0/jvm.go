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

type UndefinedAttribute struct {
    // todo
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
    _ = reader.readUint32() // attribute_length
    attrName := cp.getUtf8(attrNameIndex)
    attr := newAttribute(attrName)
    attr.readInfo(reader, cp)
    return attr
}

func newAttribute(attrName string) (AttributeInfo) {
    switch attrName {
    case "Code": return &CodeAttribute{}
    case "LineNumberTable": return &LineNumberTableAttribute{}
    case "SourceFile": return &SourceFileAttribute{}
    default: panic("BAD attr name:" + attrName) // todo
    }
}

