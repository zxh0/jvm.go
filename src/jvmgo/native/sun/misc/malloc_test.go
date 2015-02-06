package misc

import (
    "testing"
    . "jvmgo/test"
)

func TestMalloc(t *testing.T) {
    address1 := allocate(8)
    address2 := allocate(16)
    allocate(4)

    AssertEquals(8, len(memoryAt(address1)))
    AssertEquals(7, len(memoryAt(address1 + 1)))
    AssertEquals(4, len(memoryAt(address1 + 4)))

    AssertEquals(16, len(memoryAt(address2)))
    AssertEquals(13, len(memoryAt(address2 + 3)))
    AssertEquals(1, len(memoryAt(address2 + 15)))
    //memoryAt(address2 + 20)

    free(address1)
    free(address2)
}
