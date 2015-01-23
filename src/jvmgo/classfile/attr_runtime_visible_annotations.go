package classfile

/*
RuntimeVisibleAnnotations_attribute {
    u2         attribute_name_index;
    u4         attribute_length;
    u2         num_annotations;
    annotation annotations[num_annotations];
}
*/
type RuntimeVisibleAnnotationsAttribute struct {
    annotations []*Annotation
}
func (self *RuntimeVisibleAnnotationsAttribute) readInfo(reader *ClassReader, cp *ConstantPool) {
    numAnnotations := reader.readUint16()
    self.annotations = make([]*Annotation, numAnnotations)
    for i := range self.annotations {
        self.annotations[i] = readAnnotation(reader)
    }
}

/*
annotation {
    u2 type_index;
    u2 num_element_value_pairs;
    {   u2            element_name_index;
        element_value value;
    } element_value_pairs[num_element_value_pairs];
}
*/
type Annotation struct {
    typeIndex           uint16
    elementValuePairs   []*ElementValuePair
}
func readAnnotation(reader *ClassReader) (*Annotation) {
    typeIndex := reader.readUint16()
    numElementValuePairs := reader.readUint16()
    elementValuePairs := make([]*ElementValuePair, numElementValuePairs)
    for i := range elementValuePairs {
        elementValuePairs[i] = readElementValuePair(reader)
    }
    return &Annotation{typeIndex, elementValuePairs}
}

type ElementValuePair struct {
    elementNameIndex    uint16
    value               *ElementValue
}
func readElementValuePair(reader *ClassReader) (*ElementValuePair) {
    elementNameIndex := reader.readUint16()
    value := readElementValue(reader)
    return &ElementValuePair{elementNameIndex, value}
}

/*
element_value {
    u1 tag;
    union {
        u2 const_value_index;

        {   u2 type_name_index;
            u2 const_name_index;
        } enum_const_value;

        u2 class_info_index;

        annotation annotation_value;

        {   u2            num_values;
            element_value values[num_values];
        } array_value;
    } value;
}
*/
type ElementValue struct {
    tag                 uint8
    constValueIndex     uint16          // tag=B,C,D,F,I,J,S,Z,s
    enumConstValue      EnumConstValue  // tag=e
    classInfoIndex      uint16          // tag=c
    annotationValue     Annotation      // tag=@
    arrayValue          ArrayValue      // tag=[
}
func readElementValue(reader *ClassReader) (*ElementValue) {
    return nil
}

type EnumConstValue struct {
    // private U2CpIndex typeNameIndex;
    // private U2CpIndex constNameIndex;
}
type ArrayValue struct {
    // private U2 numValues;
    // private Table<ElementValue> values;
}



    // @Override
    // protected void readContent(ClassReader reader) {
    //     numValues = reader.readU2();
    //     values = reader.readTable(ElementValue.class, numValues);
    // }
    
    // @Override
    // protected void readContent(ClassReader reader) {
    //     tag = reader.readU1();
    //     tag.setDesc((char) tag.getValue());
    //     switch (tag.getValue()) {
    //         case 'B':
    //         case 'C':
    //         case 'D':
    //         case 'F':
    //         case 'I':
    //         case 'J':
    //         case 'S':
    //         case 'Z':
    //         case 's': 
    //             constValueIndex = reader.readU2CpIndex();
    //             break;
    //         case 'e': 
    //             enumConstValue = new EnumConstValue();
    //             enumConstValue.read(reader);
    //             break;
    //         case 'c':
    //             classInfoIndex = reader.readU2CpIndex();
    //             break;
    //         case '@':
    //             annotationValue = new AnnotationInfo();
    //             annotationValue.read(reader);
    //             break;
    //         case '[':
    //             arrayValue = new ArrayValue();
    //             arrayValue.read(reader);
    //             break;
    //         default: throw new FileParseException("Invalid element_value tag: " + tag.getDesc());
    //     }
    // }