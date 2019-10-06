package heap

import (
	"github.com/zxh0/jvm.go/classfile"
)

type Constant interface{}

type ConstantPool struct {
	consts []Constant
}

func newConstantPool(cf *classfile.ClassFile) ConstantPool {
	cfCp := cf.ConstantPool
	consts := make([]Constant, len(cfCp))
	rtCp := ConstantPool{consts}

	for i := 1; i < len(cfCp); i++ {
		cpInfo := cfCp[i]
		switch x := cpInfo.(type) {
		case int32, float32, string:
			consts[i] = cpInfo
		case int64, float64:
			consts[i] = cpInfo
			i++
		case classfile.ConstantStringInfo:
			consts[i] = cf.GetUTF8(x.StringIndex)
		case classfile.ConstantClassInfo:
			consts[i] = newConstantClass(cf, x)
		case classfile.ConstantFieldRefInfo:
			consts[i] = newConstantFieldRef(cf, x)
		case classfile.ConstantMethodRefInfo:
			consts[i] = newConstantMethodRef(cf, x)
		case classfile.ConstantInterfaceMethodRefInfo:
			consts[i] = newConstantInterfaceMethodRef(cf, x)
		case classfile.ConstantInvokeDynamicInfo:
			consts[i] = newConstantInvokeDynamic(cf, &rtCp, x)
		case classfile.ConstantMethodHandleInfo:
			consts[i] = newConstantMethodHandle(x)
		case classfile.ConstantMethodTypeInfo:
			consts[i] = newConstantMethodType(x)
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
