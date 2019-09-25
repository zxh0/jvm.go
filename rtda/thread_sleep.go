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

func (thread *Thread) Sleep(d time.Duration) (interrupted bool) {
	thread.lock.Lock()
	if thread.interruptedFlag {
		thread.interruptedFlag = false
		thread.lock.Unlock()
		return true
	}

	thread.sleepingFlag = true
	go thread._sleep(d)
	thread.lock.Unlock()

	interrupted = (<-thread.ch == _interrupt)
	return
}

func (thread *Thread) _sleep(d time.Duration) {
	time.Sleep(d)

	thread.lock.Lock()
	defer thread.lock.Unlock()

	if thread.sleepingFlag { // not interrupted
		thread.sleepingFlag = false
		thread.ch <- _timeOut
	}
}

func (thread *Thread) Interrupt() {
	thread.lock.Lock()
	defer thread.lock.Unlock()

	if thread.sleepingFlag {
		thread.sleepingFlag = false
		thread.ch <- _interrupt
		return
	}
	if thread.parkingFlag {
		thread.interruptedFlag = true
		thread.parkingFlag = false
		thread.ch <- _interrupt
		return
	}

	thread.interruptedFlag = true
}

func (thread *Thread) IsInterrupted(clearInterrupted bool) bool {
	thread.lock.Lock()
	defer thread.lock.Unlock()

	if thread.interruptedFlag {
		if clearInterrupted {
			thread.interruptedFlag = false
		}
		return true
	}
	return false
}

func (thread *Thread) Park(d time.Duration) {
	thread.lock.Lock()
	if thread.interruptedFlag {
		thread.interruptedFlag = false
		thread.lock.Unlock()
		return
	}
	if thread.unparkedFlag {
		thread.unparkedFlag = false
		thread.lock.Unlock()
		return
	}

	thread.parkingFlag = true
	go thread._park(d)
	thread.lock.Unlock()

	// todo: check interrupted?
	<-thread.ch
}

func (thread *Thread) _park(d time.Duration) {
	time.Sleep(d)

	thread.lock.Lock()
	defer thread.lock.Unlock()

	if thread.parkingFlag { // not interrupted
		thread.parkingFlag = false
		thread.ch <- _timeOut
	}
}

func (thread *Thread) Unpark() {
	thread.lock.Lock()
	defer thread.lock.Unlock()

	if thread.parkingFlag {
		thread.parkingFlag = false
		thread.ch <- _unpark
		return
	}

	thread.unparkedFlag = true
}
