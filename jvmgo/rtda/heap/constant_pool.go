package heap

import (
	cf "github.com/zxh0/jvm.go/jvmgo/classfile"
)

type Constant interface{}

type ConstantPool struct {
	class  *Class
	consts []Constant
}

func newConstantPool(owner *Class, cfCp *cf.ConstantPool) *ConstantPool {
	cpInfos := cfCp.Infos
	consts := make([]Constant, len(cpInfos))
	rtCp := &ConstantPool{owner, consts}

	for i := 1; i < len(cpInfos); i++ {
		cpInfo := cpInfos[i]
		switch cpInfo.(type) {
		case int32, float32, string:
			consts[i] = cpInfo
		case int64, float64:
			consts[i] = cpInfo
			i++
		case cf.ConstantStringInfo:
			strInfo := cpInfo.(cf.ConstantStringInfo)
			consts[i] = strInfo.String()
		case cf.ConstantClassInfo:
			consts[i] = newConstantClass(cpInfo.(cf.ConstantClassInfo))
		case cf.ConstantMemberrefInfo:
			consts[i] = newConstantMemberref(cpInfo.(cf.ConstantMemberrefInfo))
		case cf.ConstantInvokeDynamicInfo:
			consts[i] = newConstantInvokeDynamic(rtCp, cpInfo.(cf.ConstantInvokeDynamicInfo))
		case cf.ConstantMethodHandleInfo:
			consts[i] = newConstantMethodHandle(cpInfo.(cf.ConstantMethodHandleInfo))
		case cf.ConstantMethodTypeInfo:
			consts[i] = newConstantMethodType(cpInfo.(cf.ConstantMethodTypeInfo))
		default:
			// todo
			//fmt.Printf("%T \n", cpInfo)
			// panic("todo")
		}
	}

	return rtCp
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	// todo
	return self.consts[index]
}
