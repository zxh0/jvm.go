package classfile

var (
	_attrDeprecated = &DeprecatedAttribute{}
	_attrSynthetic  = &SyntheticAttribute{}
	_attrUndefined  = &UndefinedAttribute{}
)

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	readInfo(reader *ClassReader, attrLen uint32)
}

func readAttributes(reader *ClassReader, cp *ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp *ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrLen := reader.readUint32()
	attrName := cp.getUtf8(attrNameIndex)
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader, attrLen)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp *ConstantPool) AttributeInfo {
	switch attrName {
	case "AnnotationDefault":
		return &AnnotationDefaultAttribute{}
	// case "BootstrapMethods":
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return _attrDeprecated
	case "EnclosingMethod":
		return &EnclosingMethodAttribute{cp: cp}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "InnerClasses":
		return &InnerClassesAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "LocalVariableTypeTable":
		return &LocalVariableTypeTableAttribute{}
	// case "MethodParameters":
	case "RuntimeInvisibleAnnotations":
		return _attrUndefined
	case "RuntimeInvisibleParameterAnnotations":
		return _attrUndefined
	case "RuntimeInvisibleTypeAnnotations":
		return _attrUndefined
	case "RuntimeVisibleAnnotations":
		return &RuntimeVisibleAnnotationsAttribute{}
	case "RuntimeVisibleParameterAnnotations":
		return &ParameterAnnotationsAttribute{}
	// case "RuntimeVisibleTypeAnnotations":
	case "Signature":
		return &SignatureAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "SourceDebugExtension":
		return _attrUndefined // todo
	case "StackMapTable":
		return _attrUndefined // todo
	case "Synthetic":
		return _attrSynthetic
	default:
		return _attrUndefined
	}
}
