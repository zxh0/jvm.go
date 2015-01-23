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
    //U2CpIndex typeIndex;
    //U2 numElementValuePairs;
    //elementValuePairs;ElementValuePair
}
func readAnnotation(reader *ClassReader) (*Annotation) {
    return nil
}

type ElementValuePair struct {
    //elementNameIndex
    //value ElementValue
}


// element_value {
//     u1 tag;
//     union {
//         u2 const_value_index;

//         {   u2 type_name_index;
//             u2 const_name_index;
//         } enum_const_value;

//         u2 class_info_index;

//         annotation annotation_value;

//         {   u2            num_values;
//             element_value values[num_values];
//         } array_value;
//     } value;
// }

// type ElementValue struct {
//     private U1 tag;
//     // tag=B,C,D,F,I,J,S,Z,s
//     private U2CpIndex constValueIndex;
//     // tag=e
//     private EnumConstValue enumConstValue;
//     // tag=c
//     private U2CpIndex classInfoIndex;
//     // tag=@
//     private Annotation annotationValue;
//     // tag=[
//     private ArrayValue arrayValue;
// }
// type EnumConstValue extends ClassComponent {
//     private U2CpIndex typeNameIndex;
//     private U2CpIndex constNameIndex;
// }
// type ArrayValue extends  ClassComponent {
//     private U2 numValues;
//     private Table<ElementValue> values;
// }
    


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