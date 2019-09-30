package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type Constant interface{} // TODO: change to Slot ?

type ConstantPool struct {
	consts []Constant
}

func newConstantPool(cf *classfile.ClassFile) ConstantPool {
	cfCp := cf.ConstantPool
	cpInfos := cfCp.Infos
	consts := make([]Constant, len(cpInfos))
	rtCp := ConstantPool{consts}

	for i := 1; i < len(cpInfos); i++ {
		cpInfo := cpInfos[i]
		switch cpInfo.(type) {
		case int32, float32, string:
			consts[i] = cpInfo
		case int64, float64:
			consts[i] = cpInfo
			i++
		case classfile.ConstantStringInfo:
			strInfo := cpInfo.(classfile.ConstantStringInfo)
			consts[i] = cf.GetUTF8(strInfo.StringIndex)
		case classfile.ConstantClassInfo:
			consts[i] = newConstantClass(cf, cpInfo.(classfile.ConstantClassInfo))
		case classfile.ConstantMemberRefInfo:
			consts[i] = newConstantMemberRef(cf, cpInfo.(classfile.ConstantMemberRefInfo))
		case classfile.ConstantInvokeDynamicInfo:
			consts[i] = newConstantInvokeDynamic(cf, &rtCp, cpInfo.(classfile.ConstantInvokeDynamicInfo))
		case classfile.ConstantMethodHandleInfo:
			consts[i] = newConstantMethodHandle(cpInfo.(classfile.ConstantMethodHandleInfo))
		case classfile.ConstantMethodTypeInfo:
			consts[i] = newConstantMethodType(cpInfo.(classfile.ConstantMethodTypeInfo))
		default:
			// todo
			//fmt.Printf("%T \n", cpInfo)
			// panic("todo")
		}
	}

	return rtCp
}

func (cp ConstantPool) GetConstant(index uint) Constant {
	// todo
	return cp.consts[index]
}
