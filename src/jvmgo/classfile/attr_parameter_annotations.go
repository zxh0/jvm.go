package classfile

/*
RuntimeVisibleParameterAnnotations_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 num_parameters;
    {   u2         num_annotations;
        annotation annotations[num_annotations];
    } parameter_annotations[num_parameters];
}
RuntimeInvisibleParameterAnnotations_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 num_parameters;
    {   u2         num_annotations;
        annotation annotations[num_annotations];
    } parameter_annotations[num_parameters];
}
*/

type ParameterAnnotationsAttribute struct {
    parameterAnnotations [][]*Annotation
}

func (self *ParameterAnnotationsAttribute) readInfo(reader *ClassReader, cp *ConstantPool) {
    numParameters := reader.readUint8()
    self.parameterAnnotations = make([][]*Annotation, numParameters)

    for i := range self.parameterAnnotations {
        numAnnotations := reader.readUint16()
        self.parameterAnnotations[i] = make([]*Annotation, numAnnotations)
        
        for j := range self.parameterAnnotations[i] {
            self.parameterAnnotations[i][j] = readAnnotation(reader)
        }
    }
}
