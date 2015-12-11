package heap

import (
	"sync"
)

type Monitor struct {
	owner      interface{} // *rtda.Thread
	ownerLock  sync.Locker
	lock       sync.Locker
	entryCount int
	cond       *sync.Cond
}

func newMonitor() *Monitor {
	m := &Monitor{}
	m.ownerLock = &sync.Mutex{}
	m.lock = &sync.Mutex{}
	m.cond = sync.NewCond(m.lock)
	return m
}

func (self *Monitor) Enter(thread interface{}) {
	self.ownerLock.Lock()
	if self.owner == thread {
		self.entryCount++
		self.ownerLock.Unlock()
		return
	} else {
		self.ownerLock.Unlock()
	}

	self.lock.Lock()

	self.ownerLock.Lock()
	self.owner = thread
	self.entryCount = 1
	self.ownerLock.Unlock()
}

func (self *Monitor) Exit(thread interface{}) {
	self.ownerLock.Lock()
	var _unlock bool
	if self.owner == thread {
		self.entryCount--
		if self.entryCount == 0 {
			self.owner = nil
			_unlock = true
		}
	}
	self.ownerLock.Unlock()

	if _unlock {
		self.lock.Unlock()
	}
}

func (self *Monitor) HasOwner(thread interface{}) bool {
	self.ownerLock.Lock()
	isOwner := self.owner == thread
	self.ownerLock.Unlock()

	return isOwner
}

func (self *Monitor) Wait() {
	self.ownerLock.Lock()
	oldEntryCount := self.entryCount
	oldOwner := self.owner
	self.entryCount = 0
	self.owner = nil
	self.ownerLock.Unlock()

	self.cond.Wait()

	self.ownerLock.Lock()
	self.entryCount = oldEntryCount
	self.owner = oldOwner
	self.ownerLock.Unlock()
}

func (self *Monitor) NotifyAll() {
	self.cond.Broadcast()
}
