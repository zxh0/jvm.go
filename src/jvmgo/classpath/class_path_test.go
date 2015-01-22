package classpath

import (
    "strings"
    "testing"
)

func TestParseClassPath_Empty(t *testing.T) {
    cp := ParseClassPath("")
    if len(cp.entries) != 0 {
        t.Error("TestParseClassPath_Empty")
    }
}

func TestParseClassPath_OneDir(t *testing.T) {
    dirs := []string{".", "abc", "a/b/c"}
    for _, dir := range dirs {
        cp := ParseClassPath(dir)
        if len(cp.entries) != 1 {
            t.Error("TestParseClassPath_OneDir")
        }
    }
}

func TestParseClassPath_List(t *testing.T) {
    pathList := []string{".", "rt.jar"}
    pathStr := strings.Join(pathList, pathListSeparator)
    cp := ParseClassPath(pathStr)
    if entryCount := len(cp.entries); entryCount != 2 {
        t.Errorf("entryCount: %v", entryCount)
    }
}
