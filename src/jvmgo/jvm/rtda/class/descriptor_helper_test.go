package class

import (
	. "jvmgo/testing"
	"testing"
)

func TestCalcArgCount(t *testing.T) {
	AssertEquals(0, calcArgCount("()V"))
	AssertEquals(1, calcArgCount("(I)F"))
	AssertEquals(4, calcArgCount("([BIII)V"))
	AssertEquals(3, calcArgCount("(IDLjava/lang/Thread;)Ljava/lang/Object;"))
}

func TestGetComponentClassName(t *testing.T) {
	AssertEquals("int", getComponentClassName("[I"))
	AssertEquals("[J", getComponentClassName("[[J"))
	AssertEquals("[[D", getComponentClassName("[[[D"))
	AssertEquals("java/lang/Object", getComponentClassName("[Ljava/lang/Object;"))
	AssertEquals("[Ljava/lang/Object;", getComponentClassName("[[Ljava/lang/Object;"))
}

func TestGetReturnDescriptor(t *testing.T) {
	AssertEquals("F", GetReturnDescriptor("(I)F"))
}
