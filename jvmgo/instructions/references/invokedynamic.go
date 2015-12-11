package references

import (
	"fmt"

	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	"github.com/zxh0/jvm.go/jvmgo/rtda/heap"
)

// Bytecode Behaviors for Method Handles
const (
	REF_getField         = 1 //	getfield C.f:T
	REF_getStatic        = 2 //	getstatic C.f:T
	REF_putField         = 3 //	putfield C.f:T
	REF_putStatic        = 4 //	putstatic C.f:T
	REF_invokeVirtual    = 5 //	invokevirtual C.m:(A*)T
	REF_invokeStatic     = 6 // invokestatic C.m:(A*)T
	REF_invokeSpecial    = 7 // invokespecial C.m:(A*)T
	REF_newInvokeSpecial = 8 // new C; dup; invokespecial C.<init>:(A*)void
	REF_invokeInterface  = 9 // invokeinterface C.m:(A*)T
)

// Invoke dynamic method
type INVOKE_DYNAMIC struct {
	index uint16
	// 0
	// 0
}

func (self *INVOKE_DYNAMIC) FetchOperands(reader *base.BytecodeReader) {
	self.index = reader.ReadUint16()
	reader.ReadUint8() // must be 0
	reader.ReadUint8() // must be 0
}

func (self *INVOKE_DYNAMIC) Execute(frame *rtda.Frame) {
	self.resolveCallSiteSpecifier(frame)
	// todo
	panic("todo invokedynamic")
}

func (self *INVOKE_DYNAMIC) resolveCallSiteSpecifier(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	kIndy := cp.GetConstant(uint(self.index)).(*heap.ConstantInvokeDynamic)
	//bmSpec := kIndy.BootstrapMethodSpecifier()

	// Method Type and Method Handle Resolution

	// todo
	fmt.Printf("kIndy: %v\n", kIndy)
	kIndy.MethodHandle()
}
