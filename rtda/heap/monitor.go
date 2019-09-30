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

func (monitor *Monitor) Enter(thread interface{}) {
	monitor.ownerLock.Lock()
	if monitor.owner == thread {
		monitor.entryCount++
		monitor.ownerLock.Unlock()
		return
	} else {
		monitor.ownerLock.Unlock()
	}

	monitor.lock.Lock()

	monitor.ownerLock.Lock()
	monitor.owner = thread
	monitor.entryCount = 1
	monitor.ownerLock.Unlock()
}

func (monitor *Monitor) Exit(thread interface{}) {
	monitor.ownerLock.Lock()
	var _unlock bool
	if monitor.owner == thread {
		monitor.entryCount--
		if monitor.entryCount == 0 {
			monitor.owner = nil
			_unlock = true
		}
	}
	monitor.ownerLock.Unlock()

	if _unlock {
		monitor.lock.Unlock()
	}
}

func (monitor *Monitor) HasOwner(thread interface{}) bool {
	monitor.ownerLock.Lock()
	isOwner := monitor.owner == thread
	monitor.ownerLock.Unlock()

	return isOwner
}

func (monitor *Monitor) Wait() {
	monitor.ownerLock.Lock()
	oldEntryCount := monitor.entryCount
	oldOwner := monitor.owner
	monitor.entryCount = 0
	monitor.owner = nil
	monitor.ownerLock.Unlock()

	monitor.cond.Wait()

	monitor.ownerLock.Lock()
	monitor.entryCount = oldEntryCount
	monitor.owner = oldOwner
	monitor.ownerLock.Unlock()
}

func (monitor *Monitor) NotifyAll() {
	monitor.cond.Broadcast()
}
