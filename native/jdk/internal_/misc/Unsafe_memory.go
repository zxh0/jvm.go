package misc

import (
	"encoding/binary"
	"math"

	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_unsafe(allocateMemory, "allocateMemory", "(J)J")
	_unsafe(reallocateMemory, "reallocateMemory", "(JJ)J")
	_unsafe(freeMemory, "freeMemory", "(J)V")
	_unsafe(addressSize, "addressSize", "()I")
	_unsafe(putAddress, "putAddress", "(JJ)V")
	_unsafe(getAddress, "getAddress", "(J)J")
	_unsafe(memPutByte, "putByte", "(JB)V")
	_unsafe(memGetByte, "getByte", "(J)B")
	_unsafe(memPutShort, "putShort", "(JS)V")
	_unsafe(memGetShort, "getShort", "(J)S")
	_unsafe(memPutChar, "putChar", "(JC)V")
	_unsafe(memGetChar, "getChar", "(J)C")
	_unsafe(memPutInt, "putInt", "(JI)V")
	_unsafe(memGetInt, "getInt", "(J)I")
	_unsafe(memPutLong, "putLong", "(JJ)V")
	_unsafe(memGetLong, "getLong", "(J)J")
	_unsafe(memPutFloat, "putFloat", "(JF)V")
	_unsafe(memGetFloat, "getFloat", "(J)F")
	_unsafe(memPutDouble, "putDouble", "(JD)V")
	_unsafe(memGetDouble, "getDouble", "(J)D")
}

// public native long allocateMemory(long bytes);
// (J)J
func allocateMemory(frame *rtda.Frame) {
	// frame.GetRefVar(0) // this
	bytes := frame.GetLongVar(1)

	address := allocate(bytes)
	frame.PushLong(address)
}

// public native long reallocateMemory(long address, long bytes);
// (JJ)J
func reallocateMemory(frame *rtda.Frame) {
	// frame.GetRefVar(0) // this
	address := frame.GetLongVar(1)
	bytes := frame.GetLongVar(3)

	newAddress := reallocate(address, bytes)
	frame.PushLong(newAddress)
}

// public native void freeMemory(long address);
// (J)V
func freeMemory(frame *rtda.Frame) {
	// frame.GetRefVar(0) // this
	address := frame.GetLongVar(1)
	free(address)
}

// public native int addressSize();
// ()I
func addressSize(frame *rtda.Frame) {
	// vars := frame.
	// frame.GetRefVar(0) // this

	frame.PushInt(8) // todo unsafe.Sizeof(int)
}

// public native void putAddress(long address, long x);
// (JJ)V
func putAddress(frame *rtda.Frame) {
	memPutLong(frame)
}

// public native long getAddress(long address);
// (J)J
func getAddress(frame *rtda.Frame) {
	memGetLong(frame)
}

// public native void putByte(long address, byte x);
// (JB)V
func memPutByte(frame *rtda.Frame) {
	mem, slot := _memPut(frame)
	mem[0] = uint8(slot.IntValue())
}

// public native byte getByte(long address);
// (J)B
func memGetByte(frame *rtda.Frame) {
	mem := _memGet(frame)
	frame.PushInt(int32(mem[0]))
}

// public native void putShort(long address, short x);
// (JS)V
func memPutShort(frame *rtda.Frame) {
	mem, slot := _memPut(frame)
	binary.BigEndian.PutUint16(mem, uint16(slot.IntValue()))
}

// public native short getShort(long address);
// (J)S
func memGetShort(frame *rtda.Frame) {
	mem := _memGet(frame)
	frame.PushInt(int32(binary.BigEndian.Uint16(mem)))
}

// public native void putChar(long address, char x);
// (JC)V
func memPutChar(frame *rtda.Frame) {
	mem, slot := _memPut(frame)
	binary.BigEndian.PutUint16(mem, uint16(slot.IntValue()))
}

// public native char getChar(long address);
// (J)C
func memGetChar(frame *rtda.Frame) {
	mem := _memGet(frame)
	frame.PushInt(int32(binary.BigEndian.Uint16(mem)))
}

// public native void putInt(long address, int x);
// (JI)V
func memPutInt(frame *rtda.Frame) {
	mem, slot := _memPut(frame)
	binary.BigEndian.PutUint32(mem, uint32(slot.IntValue()))
}

// public native int getInt(long address);
// (J)I
func memGetInt(frame *rtda.Frame) {
	mem := _memGet(frame)
	frame.PushInt(int32(binary.BigEndian.Uint32(mem)))
}

// public native void putLong(long address, long x);
// (JJ)V
func memPutLong(frame *rtda.Frame) {
	mem, slot := _memPut(frame)
	binary.BigEndian.PutUint64(mem, uint64(slot.LongValue()))
}

// public native long getLong(long address);
// (J)J
func memGetLong(frame *rtda.Frame) {
	mem := _memGet(frame)
	frame.PushLong(int64(binary.BigEndian.Uint64(mem)))
}

// public native void putFloat(long address, float x);
// (JJ)V
func memPutFloat(frame *rtda.Frame) {
	mem, slot := _memPut(frame)
	binary.BigEndian.PutUint32(mem, uint32(slot.IntValue()))
}

// public native float getFloat(long address);
// (J)J
func memGetFloat(frame *rtda.Frame) {
	mem := _memGet(frame)
	frame.PushFloat(math.Float32frombits(binary.BigEndian.Uint32(mem)))
}

// public native void putDouble(long address, double x);
// (JJ)V
func memPutDouble(frame *rtda.Frame) {
	mem, slot := _memPut(frame)
	binary.BigEndian.PutUint64(mem, uint64(slot.LongValue()))
}

// public native double getDouble(long address);
// (J)J
func memGetDouble(frame *rtda.Frame) {
	mem := _memGet(frame)
	frame.PushDouble(math.Float64frombits(binary.BigEndian.Uint64(mem)))
}

func _memPut(frame *rtda.Frame) ([]byte, heap.Slot) {
	// frame.GetRefVar(0) // this
	address := frame.GetLongVar(1)
	slot := frame.GetLocalVar(3)
	mem := memoryAt(address)
	return mem, slot
}

func _memGet(frame *rtda.Frame) []byte {
	// frame.GetRefVar(0) // this
	address := frame.GetLongVar(1)
	return memoryAt(address)
}
