package misc

import "jvmgo/util"

var _allocated = map[int64][]byte{}
var _nextAddress = int64(64) // not zero!

func allocate(size int64) (address int64) {
    mem := make([]byte, size)
    address = _nextAddress
    _allocated[address] = mem
    _nextAddress += size
    return 
}

func free(address int64) {
    if _, ok := _allocated[address]; ok {
        delete(_allocated, address)
    } else {
        util.Panicf("memory was not allocated: %v", address)
    }
}

func memoryAt(address int64) ([]byte) {
    for startAddress, mem := range _allocated {
        endAddress := startAddress + int64(len(mem))
        if address >= startAddress && address < endAddress {
            offset := address - startAddress
            return mem[offset:]
        }
    }
    util.Panicf("invalid address: %v", address)
    return nil
}
