package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/linker"
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
	cp := frame.GetConstantPool()
	ref := cp.GetConstantInvokeDynamic(uint(instr.index))
	linker.ResolveInvokeDynamic(frame.GetBootLoader(), ref)
	// TODO
	println(ref.Name, ref.Type)
	println(ref.MethodRef.ResolvedClass.Name, ref.MethodRef.ResolvedMethod.Name)
	panic("TODO: invokedynamic")
}
