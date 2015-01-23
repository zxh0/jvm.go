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

func newConstantPool(class *Class, cfCp *cf.ConstantPool) (*ConstantPool) {
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
        // todo 
        }
    }

    return rtCp
}
