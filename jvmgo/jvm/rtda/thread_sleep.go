package rtda

import (
	"time"
)

const (
	// wake up signals
	_timeOut   = 1
	_interrupt = 2
	_unpark    = 3
)

func (self *Thread) Sleep(d time.Duration) (interrupted bool) {
	self.lock.Lock()
	if self.interruptedFlag {
		self.interruptedFlag = false
		self.lock.Unlock()
		return true
	}

	self.blockingFlag = true
	go self._sleep(d)
	self.lock.Unlock()

	interrupted = (<-self.ch == _interrupt)
	return
}

func (self *Thread) _sleep(d time.Duration) {
	time.Sleep(d)

	self.lock.Lock()
	defer self.lock.Unlock()

	if self.blockingFlag { // not interrupted
		self.blockingFlag = false
		self.ch <- _timeOut
	}
}

func (self *Thread) Interrupt() {
	self.lock.Lock()
	defer self.lock.Unlock()

	if self.blockingFlag {
		self.blockingFlag = false
		self.ch <- _interrupt
		return
	}

	self.interruptedFlag = true
}

func (self *Thread) IsInterrupted(clearInterrupted bool) bool {
	self.lock.Lock()
	defer self.lock.Unlock()

	if self.interruptedFlag {
		if clearInterrupted {
			self.interruptedFlag = false
		}
		return true
	}
	return false
}

func (self *Thread) Park(d time.Duration) {
	self.Sleep(d)
}

func (self *Thread) Unpark() {

}
