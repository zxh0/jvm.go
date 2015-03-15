package classpath

import (
	. "jvmgo/testing"
	"strings"
	"testing"
)

func TestParseClassPath_Empty(t *testing.T) {
	cp := ParseClassPath("").compoundEntry
	AssertEquals(2, len(cp.entries))
}

func TestParseClassPath_OneDir(t *testing.T) {
	dirs := []string{".", "abc", "a/b/c"}
	for _, dir := range dirs {
		cp := ParseClassPath(dir).compoundEntry
		AssertEquals(2, len(cp.entries))
	}
}

func TestParseClassPath_List(t *testing.T) {
	pathList := []string{".", "rt.jar"}
	pathStr := strings.Join(pathList, _pathListSeparator)
	cp := ParseClassPath(pathStr).compoundEntry
	AssertEquals(2, len(cp.entries))
}
