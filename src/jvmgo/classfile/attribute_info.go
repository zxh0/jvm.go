package classfile

import "fmt"

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
    // todo
}

type UndefinedAttribute struct {
    // todo
}

func readAttribute(reader *ClassReader, cp *ConstantPool) (AttributeInfo) {
    attrNameIndex := reader.readUint16()
    _ = reader.readUint32() // attribute_length
    attrName := cp.getUtf8(attrNameIndex)
    
    switch attrName {
    case "Code": return readCodeAttribute(reader, cp)
    case "LineNumberTable": return readLineNumberTableAttribute(reader)
    case "SourceFile": return readSourceFileAttribute(reader)
    //default:
    }

    fmt.Printf("attrName: %v \n", attrName)
    // todo
    panic("todo attr:" + attrName)
    return nil
}
