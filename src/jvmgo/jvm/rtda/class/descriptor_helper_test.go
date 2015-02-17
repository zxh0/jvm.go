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

func TestGetComponentDescriptor(t *testing.T) {
	AssertEquals("I", getComponentDescriptor("[I"))
	AssertEquals("[J", getComponentDescriptor("[[J"))
	AssertEquals("Ljava/lang/Object;", getComponentDescriptor("[Ljava/lang/Object;"))
}

func TestGetClassName(t *testing.T) {
	AssertEquals("double", getClassName("D"))
	AssertEquals("java/lang/Object", getClassName("Ljava/lang/Object;"))
	AssertEquals("[F", getClassName("[F"))
	AssertEquals("[[B", getClassName("[[B"))
}

func TestGetReturnDescriptor(t *testing.T) {
	AssertEquals("F", GetReturnDescriptor("(I)F"))
}
