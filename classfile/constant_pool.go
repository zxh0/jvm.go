package classfile

import (
	"fmt"
)

type ConstantPool struct {
	cf    *ClassFile
	Infos []ConstantInfo
}

func (cp *ConstantPool) read(reader *ClassReader) {
	cpCount := int(reader.readUint16())
	cp.Infos = make([]ConstantInfo, cpCount)

	// The constant_pool table is indexed from 1 to constant_pool_count - 1.
	for i := 1; i < cpCount; i++ {
		cp.Infos[i] = readConstantInfo(reader, cp)
		// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.5
		// All 8-byte constants take up two entries in the constant_pool table of the class file.
		// If a CONSTANT_Long_info or CONSTANT_Double_info structure is the item in the constant_pool
		// table at index n, then the next usable item in the pool is located at index n+2.
		// The constant_pool index n+1 must be valid but is considered unusable.
		switch cp.Infos[i].(type) {
		case int64, float64:
			i++
		}
	}
}

func (cp *ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	cpInfo := cp.Infos[index]
	if cpInfo == nil {
		panic(fmt.Errorf("Bad constant pool index: %v!", index))
	}

	return cpInfo
}

func (cp *ConstantPool) getNameAndType(index uint16) (name, _type string) {
	ntInfo := cp.getConstantInfo(index).(ConstantNameAndTypeInfo)
	name = cp.getUtf8(ntInfo.nameIndex)
	_type = cp.getUtf8(ntInfo.descriptorIndex)
	return
}

func (cp *ConstantPool) getClassName(index uint16) string {
	classInfo := cp.getConstantInfo(index).(ConstantClassInfo)
	return cp.getUtf8(classInfo.nameIndex)
}

func (cp *ConstantPool) getUtf8(index uint16) string {
	return cp.getConstantInfo(index).(string)
}
