package class

import (
    "sync"
    . "jvmgo/any"
)

type Monitor struct {
    owner       Any // *rtda.Thread
    ownerLock   sync.Mutex
    lock        sync.Mutex
    entryCount  int
}

func (self *Monitor) Enter(thread Any) {
    self.ownerLock.Lock()
    if self.owner == thread {
        self.entryCount++
        self.ownerLock.Unlock()
        return
    } else {
        self.ownerLock.Unlock()
    }

    self.lock.Lock()
    self.owner = thread
    self.entryCount = 1
}

func (self *Monitor) Exit(thread Any) {
    self.ownerLock.Lock()
    if self.owner == thread {
        self.entryCount--
        if self.entryCount == 0 {
            self.owner = nil
            self.lock.Unlock()
        }
    }
    self.ownerLock.Unlock()
}
