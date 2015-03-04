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
	UndefinedAttribute
}

/*
AnnotationDefault_attribute {
    u2            attribute_name_index;
    u4            attribute_length;
    element_value default_value;
}
*/
type AnnotationDefaultAttribute struct {
	UndefinedAttribute
}
