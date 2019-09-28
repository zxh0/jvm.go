package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

func getNameAndType(cf *classfile.ClassFile, index uint16) (name, _type string) {
	if index > 0 {
		ntInfo := cf.GetConstantInfo(index).(classfile.ConstantNameAndTypeInfo)
		name = cf.GetUTF8(ntInfo.NameIndex)
		_type = cf.GetUTF8(ntInfo.DescriptorIndex)
	}
	return
}
