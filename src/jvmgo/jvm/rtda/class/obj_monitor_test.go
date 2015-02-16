package class

import (
    "testing"
    . "jvmgo/test"
)

func TestMonitor(t *testing.T) {
    thread := "not thread!"

    monitor := &Monitor{}
    AssertEquals(0, monitor.entryCount)
    monitor.Enter(thread)
    AssertEquals(1, monitor.entryCount)
    monitor.Enter(thread)
    AssertEquals(2, monitor.entryCount)

    monitor.Exit(thread)
    AssertEquals(1, monitor.entryCount)
    monitor.Exit(thread)
    AssertEquals(0, monitor.entryCount)
}
