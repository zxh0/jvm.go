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
    for i := 1; i < len(cpInfos); i++ {
        cpInfo := cpInfos[i]
        switch cpInfo.(type) {
        case *cf.ConstantIntegerInfo:
            cInt := cpInfo.(*cf.ConstantIntegerInfo)
            consts[i] = cInt.Value()
        case *cf.ConstantFloatInfo:
            cFloat := cpInfo.(*cf.ConstantFloatInfo)
            consts[i] = cFloat.Value()
        case *cf.ConstantLongInfo:
            cLong := cpInfo.(*cf.ConstantLongInfo)
            consts[i] = cLong.Value()
            i++
        case *cf.ConstantDoubleInfo:
            cDouble := cpInfo.(*cf.ConstantDoubleInfo)
            consts[i] = cDouble.Value()
            i++
            //case *cf.
        // todo methodref
        }
    }

    // todo
    return &ConstantPool{nil, consts}
}
