package class

import (
    "sync"
    . "jvmgo/any"
)

type ReentrantLock struct {
    holder      Any // *rtda.Thread
    holderLock  sync.Mutex
    lock        sync.Mutex
    lockCount   int
}

func (self *ReentrantLock) Lock(thread Any) {
    self.holderLock.Lock()
    if self.holder == thread {
        self.lockCount++
        self.holderLock.Unlock()
        return
    } else {
        self.holderLock.Unlock()
    }

    self.lock.Lock()
    self.holder = thread
    self.lockCount = 1
}

func (self *ReentrantLock) Unlock(thread Any) {
    self.holderLock.Lock()
    if self.holder == thread {
        self.lockCount--
        if self.lockCount == 0 {
            self.holder = nil
            self.lock.Unlock()
        }
    }
    self.holderLock.Unlock()
}
