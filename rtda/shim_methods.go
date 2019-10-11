package rtda

import (
	"github.com/zxh0/jvm.go/rtda/heap"
)

var (
	_shimClass = &heap.Class{Name: "~shim"}

	shimReturnMethod = &heap.Method{
		ClassMember: heap.ClassMember{
			AccessFlags: heap.AccStatic,
			Name:        "<return>",
			Class:       _shimClass,
		},
		MethodData: heap.MethodData{
			Code: []byte{0xb1}, // return
		},
	}

	shimAThrowMethod = &heap.Method{
		ClassMember: heap.ClassMember{
			AccessFlags: heap.AccStatic,
			Name:        "<athrow>",
			Class:       _shimClass,
		},
		MethodData: heap.MethodData{
			Code: []byte{0xbf}, // athrow
		},
	}

	ShimBootstrapMethod = &heap.Method{
		ClassMember: heap.ClassMember{
			AccessFlags: heap.AccStatic,
			Name:        "<bootstrap>",
			Class:       _shimClass,
		},
		MethodData: heap.MethodData{
			MaxStack:  8,
			MaxLocals: 8,
			Code:      []byte{0xff, 0xb1}, // bootstrap, return
		},
		ParamSlotCount: 2,
	}
)
