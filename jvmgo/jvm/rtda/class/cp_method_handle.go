package class

import (
	cf "github.com/zxh0/jvm.go/jvmgo/classfile"
)

type ConstantMethodHandle struct {
	referenceKind  uint8
	referenceIndex uint16
}

func newConstantMethodHandle(cp *ConstantPool, mhInfo *cf.ConstantMethodHandleInfo) *ConstantMethodHandle {
	return &ConstantMethodHandle{
		referenceKind:  mhInfo.ReferenceKind(),
		referenceIndex: mhInfo.ReferenceIndex(),
	}
}
