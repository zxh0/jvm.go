package references

import (
	"fmt"

	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

// Bytecode Behaviors for Method Handles
const (
	RefGetField         = 1 // getfield C.f:T
	RefGetStatic        = 2 // getstatic C.f:T
	RefPutField         = 3 // putfield C.f:T
	RefPutStatic        = 4 // putstatic C.f:T
	RefInvokeVirtual    = 5 // invokevirtual C.m:(A*)T
	RefInvokeStatic     = 6 // invokestatic C.m:(A*)T
	RefInvokeSpecial    = 7 // invokespecial C.m:(A*)T
	RefNewInvokeSpecial = 8 // new C; dup; invokespecial C.<init>:(A*)void
	RefInvokeInterface  = 9 // invokeinterface C.m:(A*)T
)

// Invoke dynamic method
type InvokeDynamic struct {
	index uint16
	// 0
	// 0
}

func (instr *InvokeDynamic) FetchOperands(reader *base.CodeReader) {
	instr.index = reader.ReadUint16()
	reader.ReadUint8() // must be 0
	reader.ReadUint8() // must be 0
}

func (instr *InvokeDynamic) Execute(frame *rtda.Frame) {
	instr.resolveCallSiteSpecifier(frame)
	// todo
	panic("todo invokedynamic")
}

func (instr *InvokeDynamic) resolveCallSiteSpecifier(frame *rtda.Frame) {
	cp := frame.GetConstantPool()
	kIndy := cp.GetConstant(uint(instr.index)).(*heap.ConstantInvokeDynamic)
	//bmSpec := kIndy.BootstrapMethodSpecifier()

	// Method Type and Method Handle Resolution

	// todo
	fmt.Printf("kIndy: %v\n", kIndy)
	kIndy.MethodHandle()
}
