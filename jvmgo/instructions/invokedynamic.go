package instructions

import (
	"fmt"

	"github.com/zxh0/jvm.go/jvmgo/jvm/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
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
type invokedynamic struct {
	index uint16
	// 0
	// 0
}

func (self *invokedynamic) fetchOperands(decoder *InstructionDecoder) {
	self.index = decoder.readUint16()
	decoder.readUint8() // must be 0
	decoder.readUint8() // must be 0
}

func (self *invokedynamic) Execute(frame *rtda.Frame) {
	self.resolveCallSiteSpecifier(frame)
	// todo
	panic("todo invokedynamic")
}

func (self *invokedynamic) resolveCallSiteSpecifier(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	kIndy := cp.GetConstant(uint(self.index)).(*rtc.ConstantInvokeDynamic)
	//bmSpec := kIndy.BootstrapMethodSpecifier()

	// Method Type and Method Handle Resolution

	// todo
	fmt.Printf("kIndy: %v\n", kIndy)
	kIndy.MethodHandle()
}
