package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type ClassMember struct {
	classfile.AccessFlags
	Name           string
	Descriptor     string
	Signature      string
	AnnotationData []byte // RuntimeVisibleAnnotations_attribute
	Class          *Class
}

func (m *ClassMember) copyMemberData(cf *classfile.ClassFile, cfMember classfile.MemberInfo) {
	m.AccessFlags = classfile.AccessFlags(cfMember.AccessFlags)
	m.Name = cf.GetUTF8(cfMember.NameIndex)
	m.Descriptor = cf.GetUTF8(cfMember.DescriptorIndex)
	m.Signature = cf.GetUTF8(cfMember.GetSignatureIndex())
	m.AnnotationData = cfMember.GetRuntimeVisibleAnnotationsAttributeData()
}
