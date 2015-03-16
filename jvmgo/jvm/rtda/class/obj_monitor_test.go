package class

import (
	. "github.com/zxh0/jvm.go/jvmgo/testing"
	"testing"
)

func TestMonitor(t *testing.T) {
	thread := "not thread!"

	monitor := newMonitor()
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
