package class

import (
    "sync"
    . "jvmgo/any"
    "jvmgo/util"
)

type ReentrantLock struct {
    holder      Any // *rtda.Thread
    holderLock  sync.Mutex
    lock        sync.Mutex
    lockCount   int
}

// thread: *rtda.Thread
func (self *ReentrantLock) Lock(thread Any) {
    self.holderLock.Lock()
    defer self.holderLock.Unlock()

    if self.holder == thread {
        self.lockCount++
    } else if self.holder == nil {
        self.holder = thread
        self.lockCount++
        self.lock.Lock()
    }
}

func (self *ReentrantLock) Unlock(thread Any) {
    self.holderLock.Lock()
    defer self.holderLock.Unlock()

    if self.holder == thread && self.lockCount > 0 {
        self.lockCount--
        if self.lockCount == 0 {
            self.holder = nil
            self.lock.Unlock()
        }
    } else {
        // todo
        util.Panicf("BAD ReentrantLock state! holder:%v lockCount:%v",
            self.holder, self.lockCount)
    }
}
