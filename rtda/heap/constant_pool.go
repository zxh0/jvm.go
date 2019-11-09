package heap

import (
	"github.com/zxh0/jvm.go/classfile"
	"github.com/zxh0/jvm.go/vmutils"
)

/*
symbolic_ref
  class_ref
  member_ref
    field_ref
    method_ref
    interface_method_ref
  method_type_ref
  method_handle_ref
*/

type Constant interface{}
type ConstantPool []Constant

func newConstantPool(class *Class, cf *classfile.ClassFile) ConstantPool {
	cfCp := cf.ConstantPool
	rtCp := make([]Constant, len(cfCp))

	for i := 1; i < len(cfCp); i++ {
		cpInfo := cfCp[i]
		switch x := cpInfo.(type) {
		case int32, float32:
			rtCp[i] = cpInfo
		case int64, float64:
			rtCp[i] = cpInfo
			i++
		case []byte: // utf8
			rtCp[i] = vmutils.DecodeMUTF8(x)
		case classfile.ConstantStringInfo:
			rtCp[i] = newConstantString(class, cf.GetUTF8(x.StringIndex))
		case classfile.ConstantClassInfo:
			rtCp[i] = newConstantClass(cf, x)
		case classfile.ConstantFieldRefInfo:
			rtCp[i] = newConstantFieldRef(cf, x)
		case classfile.ConstantMethodRefInfo:
			rtCp[i] = newConstantMethodRef(cf, x)
		case classfile.ConstantInvokeDynamicInfo:
			rtCp[i] = newConstantInvokeDynamic(cf, x)
		case classfile.ConstantMethodHandleInfo:
			rtCp[i] = newConstantMethodHandle(x)
		case classfile.ConstantMethodTypeInfo:
			rtCp[i] = newConstantMethodType(x)
		default:
			// todo
			//fmt.Printf("%T \n", cpInfo)
			// panic("todo")
		}
	}

	return rtCp
}

func (cp ConstantPool) GetConstantString(index uint) *ConstantString {
	return cp.GetConstant(index).(*ConstantString)
}
func (cp ConstantPool) GetConstantClass(index uint) *ConstantClass {
	return cp.GetConstant(index).(*ConstantClass)
}
func (cp ConstantPool) GetConstantFieldRef(index uint) *ConstantFieldRef {
	return cp.GetConstant(index).(*ConstantFieldRef)
}
func (cp ConstantPool) GetConstantInvokeDynamic(index uint) *ConstantInvokeDynamic {
	return cp.GetConstant(index).(*ConstantInvokeDynamic)
}

func (cp ConstantPool) GetConstant(index uint) Constant {
	// TODO: check index
	return cp[index]
}
