package class

import cf "jvmgo/classfile"

type Constant interface{}

type ConstantPool struct {
    class  *Class
    consts []Constant
}

// func (self *ConstantPool) GetFieldref(index uint) (*ConstantFieldref) {
//     return self.GetConstant()
// }

func (self *ConstantPool) GetConstant(index uint) (Constant) {
    // todo
    return self.consts[index]
}

func newConstantPool(cfCp *cf.ConstantPool) (*ConstantPool) {
    cpInfos := cfCp.Infos()
    consts := make([]Constant, len(cpInfos))
    rtCp := &ConstantPool{nil, consts} // todo

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
        case *cf.ConstantClassInfo:
            classInfo := cpInfo.(*cf.ConstantClassInfo)
            cClass := &ConstantClass{}
            cClass.cp = rtCp
            cClass.name = classInfo.Name()
            consts[i] = cClass
        case *cf.ConstantFieldrefInfo:
        case *cf.ConstantMethodrefInfo:
        case *cf.ConstantInterfaceMethodrefInfo:
        // todo methodref
        }
    }

    return rtCp
}
