package references

import (
	"github.com/zxh0/jvm.go/jvmgo/instructions/base"
	"github.com/zxh0/jvm.go/jvmgo/rtda"
	rtc "github.com/zxh0/jvm.go/jvmgo/rtda/class"
)

// Create new array
type NEW_ARRAY struct {
	atype uint8
}

func (self *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.atype = reader.ReadUint8()
}
func (self *NEW_ARRAY) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		frame.Thread().ThrowNegativeArraySizeException()
		return
	}

	arr := rtc.NewPrimitiveArray(self.atype, uint(count))
	stack.PushRef(arr)
}

// Create new array of reference
type ANEW_ARRAY struct{ base.Index16Instruction }

func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.ConstantPool()
	kClass := cp.GetConstant(self.Index).(*rtc.ConstantClass)
	componentClass := kClass.Class()

	if componentClass.InitializationNotStarted() {
		thread := frame.Thread()
		frame.SetNextPC(thread.PC()) // undo anewarray
		thread.InitClass(componentClass)
		return
	}

	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		frame.Thread().ThrowNegativeArraySizeException()
	} else {
		arr := rtc.NewRefArray(componentClass, uint(count))
		stack.PushRef(arr)
	}

}

// Create new multidimensional array
type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
}

func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.index = reader.ReadUint16()
	self.dimensions = reader.ReadUint8()
}
func (self *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.ConstantPool()
	kClass := cp.GetConstant(uint(self.index)).(*rtc.ConstantClass)
	arrClass := kClass.Class()

	stack := frame.OperandStack()
	counts := stack.PopTops(uint(self.dimensions))
	if !_checkCounts(counts) {
		frame.Thread().ThrowNegativeArraySizeException()
	} else {
		arr := _newMultiArray(counts, arrClass)
		stack.PushRef(arr)
	}
}

func _checkCounts(counts []interface{}) bool {
	for _, c := range counts {
		if c.(int32) < 0 {
			return false
		}
	}
	return true
}

func _newMultiArray(counts []interface{}, arrClass *rtc.Class) *rtc.Obj {
	count := uint(counts[0].(int32))
	arr := rtc.NewArray(arrClass, count)

	if len(counts) > 1 {
		objs := arr.Refs()
		for i := range objs {
			objs[i] = _newMultiArray(counts[1:], arrClass.ComponentClass())
		}
	}

	return arr
}
