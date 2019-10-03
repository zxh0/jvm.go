package classfile

import (
	"fmt"
)

type ConstantPool struct {
	Infos []ConstantInfo
}

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.ReadUint16())
	consts := make([]ConstantInfo, cpCount)

	// The constant_pool table is indexed from 1 to constant_pool_count - 1.
	for i := 1; i < cpCount; i++ {
		consts[i] = readConstantInfo(reader)
		// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.5
		// All 8-byte constants take up two entries in the constant_pool table of the class file.
		// If a CONSTANT_Long_info or CONSTANT_Double_info structure is the item in the constant_pool
		// table at index n, then the next usable item in the pool is located at index n+2.
		// The constant_pool index n+1 must be valid but is considered unusable.
		switch consts[i].(type) {
		case int64, float64:
			i++
		}
	}

	return ConstantPool{Infos: consts}
}

func (cp *ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := cp.Infos[index]; cpInfo == nil {
		panic(fmt.Errorf("invalid constant pool index: %d", index))
	} else {
		return cpInfo
	}
}

func (cp *ConstantPool) getUtf8(index uint16) string {
	if index == 0 {
		return ""
	}
	return cp.getConstantInfo(index).(string)
}
