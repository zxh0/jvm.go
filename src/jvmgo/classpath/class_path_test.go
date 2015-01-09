package classpath

import "testing"

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
