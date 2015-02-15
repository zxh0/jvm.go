package keepalive

import (
    "sync"
)

var (
    aliveCount = 0
    lock = &sync.Mutex{}
    cond = sync.NewCond(lock)
)

func NonDaemonThreadStart() {
    lock.Lock()
    defer lock.Unlock()
    
    aliveCount++
}

func NonDaemonThreadStop() {
    lock.Lock()
    defer lock.Unlock()

    aliveCount--
    if aliveCount == 0 {
        cond.Broadcast()
    }
}

func KeepAlive() {
    lock.Lock()
    defer lock.Unlock()

    if aliveCount > 0 {
        cond.Wait()
    }
}
