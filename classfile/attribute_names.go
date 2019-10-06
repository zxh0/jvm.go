package classfile

// Predefined class file attributes
const (
	ConstantValue                       = "ConstantValue"                        //	1.0.2
	Code                                = "Code"                                 //	1.0.2
	Exceptions                          = "Exceptions"                           //	1.0.2
	SourceFile                          = "SourceFile"                           //	1.0.2
	LineNumberTable                     = "LineNumberTable"                      //	1.0.2
	LocalVariableTable                  = "LocalVariableTable"                   //	1.0.2
	InnerClasses                        = "InnerClasses"                         //	1.1
	Synthetic                           = "Synthetic"                            //	1.1
	Deprecated                          = "Deprecated"                           //	1.1
	EnclosingMethod                     = "EnclosingMethod"                      //	5.0
	Signature                           = "Signature"                            //	5.0
	SourceDebugExtension                = "SourceDebugExtension"                 //	5.0
	LocalVariableTypeTable              = "LocalVariableTypeTable"               //	5.0
	RuntimeVisibleAnnotations           = "RuntimeVisibleAnnotations"            //	5.0
	RuntimeInvisibleAnnotations         = "RuntimeInvisibleAnnotations"          //	5.0
	RuntimeVisibleParameterAnnotations  = "RuntimeVisibleParameterAnnotations"   //	5.0
	RuntimeInvisibleParameterAnnotation = "RuntimeInvisibleParameterAnnotations" //	5.0
	AnnotationDefault                   = "AnnotationDefault"                    //	5.0
	StackMapTable                       = "StackMapTable"                        //	6
	BootstrapMethods                    = "BootstrapMethods"                     //	7
	RuntimeVisibleTypeAnnotations       = "RuntimeVisibleTypeAnnotations"        //	8
	RuntimeInvisibleTypeAnnotations     = "RuntimeInvisibleTypeAnnotations"      //	8
	MethodParameters                    = "MethodParameters"                     //	8
	Module                              = "Module"                               // 9
	ModulePackages                      = "ModulePackages"                       // 9
	ModuleMainClass                     = "ModuleMainClass"                      // 9
	NestHost                            = "NestHost"                             // 11
	NestMembers                         = "NestMembers"                          // 11
)
