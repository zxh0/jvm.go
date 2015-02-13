package class

import (
    "testing"
    . "jvmgo/test"
)

func TestReentrantLock(t *testing.T) {
    holder := "x"

    lock := &ReentrantLock{}
    AssertEquals(0, lock.lockCount)
    lock.Lock(holder)
    AssertEquals(1, lock.lockCount)
    lock.Lock(holder)
    AssertEquals(2, lock.lockCount)

    lock.Unlock(holder)
    AssertEquals(1, lock.lockCount)
    lock.Unlock(holder)
    AssertEquals(0, lock.lockCount)
}
