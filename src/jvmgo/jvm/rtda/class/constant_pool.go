package class

import (
	//"fmt"
	cf "jvmgo/classfile"
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
			intInfo := cpInfo.(*cf.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *cf.ConstantFloatInfo:
			floatInfo := cpInfo.(*cf.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *cf.ConstantLongInfo:
			longInfo := cpInfo.(*cf.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *cf.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*cf.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *cf.ConstantStringInfo:
			stringInfo := cpInfo.(*cf.ConstantStringInfo)
			consts[i] = newConstantString(stringInfo)
		case *cf.ConstantClassInfo:
			classInfo := cpInfo.(*cf.ConstantClassInfo)
			consts[i] = newConstantClass(rtCp, classInfo)
		case *cf.ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*cf.ConstantFieldrefInfo)
			consts[i] = newConstantFieldref(rtCp, fieldrefInfo)
		case *cf.ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*cf.ConstantMethodrefInfo)
			consts[i] = newConstantMethodref(rtCp, methodrefInfo)
		case *cf.ConstantInterfaceMethodrefInfo:
			methodrefInfo := cpInfo.(*cf.ConstantInterfaceMethodrefInfo)
			consts[i] = newConstantInterfaceMethodref(rtCp, methodrefInfo)
		case *cf.ConstantUtf8Info:
			utf8Info := cpInfo.(*cf.ConstantUtf8Info)
			consts[i] = newConstantUtf8(utf8Info)
		default:
			// todo
			//fmt.Printf("%v \n", cpInfo)
			//panic("todo")
		}
	}

	return rtCp
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	// todo
	return self.consts[index]
}

// todo
func (self *ConstantPool) GetMethodref(name string) *ConstantMethodref {
	for _, c := range self.consts {
		if methodref, ok := c.(*ConstantMethodref); ok {
			if methodref.name == name {
				return methodref
			}
		}
	}
	return nil // todo
}
