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
	cpInfos := cfCp.Infos()
	consts := make([]Constant, len(cpInfos))
	rtCp := &ConstantPool{owner, consts}

	for i := 1; i < len(cpInfos); i++ {
		cpInfo := cpInfos[i]
		switch cpInfo.(type) {
		case *cf.ConstantIntegerInfo:
			consts[i] = cpInfo.(*cf.ConstantIntegerInfo).Value()
		case *cf.ConstantFloatInfo:
			consts[i] = cpInfo.(*cf.ConstantFloatInfo).Value()
		case *cf.ConstantLongInfo:
			consts[i] = cpInfo.(*cf.ConstantLongInfo).Value()
			i++
		case *cf.ConstantDoubleInfo:
			consts[i] = cpInfo.(*cf.ConstantDoubleInfo).Value()
			i++
		case *cf.ConstantStringInfo:
			consts[i] = cpInfo.(*cf.ConstantStringInfo).String()
		case *cf.ConstantUtf8Info:
			consts[i] = newConstantUtf8(cpInfo.(*cf.ConstantUtf8Info))
		case *cf.ConstantClassInfo:
			consts[i] = newConstantClass(cpInfo.(*cf.ConstantClassInfo))
		case *cf.ConstantFieldrefInfo:
			consts[i] = newConstantFieldref(cpInfo.(*cf.ConstantFieldrefInfo))
		case *cf.ConstantMethodrefInfo:
			consts[i] = newConstantMethodref(cpInfo.(*cf.ConstantMethodrefInfo))
		case *cf.ConstantInterfaceMethodrefInfo:
			consts[i] = newConstantInterfaceMethodref(cpInfo.(*cf.ConstantInterfaceMethodrefInfo))
		case *cf.ConstantInvokeDynamicInfo:
			consts[i] = newConstantInvokeDynamic(rtCp, cpInfo.(*cf.ConstantInvokeDynamicInfo))
		case *cf.ConstantMethodHandleInfo:
			consts[i] = newConstantMethodHandle(cpInfo.(*cf.ConstantMethodHandleInfo))
		case *cf.ConstantMethodTypeInfo:
			consts[i] = newConstantMethodType(cpInfo.(*cf.ConstantMethodTypeInfo))
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
