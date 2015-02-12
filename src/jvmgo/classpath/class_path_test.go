package classpath

import (
    "strings"
    "testing"
    . "jvmgo/test"
)

func TestParseClassPath_Empty(t *testing.T) {
    cp := ParseClassPath("")
    AssertEquals(0, len(cp.entries))
}

func TestParseClassPath_OneDir(t *testing.T) {
    dirs := []string{".", "abc", "a/b/c"}
    for _, dir := range dirs {
        cp := ParseClassPath(dir)
        AssertEquals(1, len(cp.entries))
    }
}

func TestParseClassPath_List(t *testing.T) {
    pathList := []string{".", "rt.jar"}
    pathStr := strings.Join(pathList, pathListSeparator)
    cp := ParseClassPath(pathStr)
    AssertEquals(2, len(cp.entries))
}
