package misc

import (
	"github.com/zxh0/jvm.go/jvmgo/jutil"
)

var _allocated = map[int64][]byte{}
var _nextAddress = int64(64) // not zero!

func allocate(size int64) (address int64) {
	mem := make([]byte, size)
	address = _nextAddress
	_allocated[address] = mem
	_nextAddress += size
	return
}

func reallocate(address, size int64) int64 {
	if size == 0 {
		return 0
	} else if address == 0 {
		return allocate(size)
	} else {
		mem := memoryAt(address)
		if len(mem) >= int(size) {
			return address
		} else {
			delete(_allocated, address)
			newAddress := allocate(size)
			newMem := memoryAt(newAddress)
			copy(newMem, mem)
			return newAddress
		}
	}
}

func free(address int64) {
	if _, ok := _allocated[address]; ok {
		delete(_allocated, address)
	} else {
		jutil.Panicf("memory was not allocated: %v", address)
	}
}

func memoryAt(address int64) []byte {
	for startAddress, mem := range _allocated {
		endAddress := startAddress + int64(len(mem))
		if address >= startAddress && address < endAddress {
			offset := address - startAddress
			return mem[offset:]
		}
	}
	jutil.Panicf("invalid address: %v", address)
	return nil
}
