package test

import "testing"
import "jvmgo/classpath"

func TestParseClassPath_Empty(t *testing.T) {
    cp := classpath.ParseClassPath("")
    if len(cp.entries) != 0 {
        t.Error("TestParseClassPath_Empty")
    }
}

func TestParseClassPath_OneDir(t *testing.T) {
    dirs := []string{".", "abc", "a/b/c"}
    for _, dir := range dirs {
        cp := classpath.ParseClassPath(dir)
        if len(cp.entries) != 1 {
            t.Error("TestParseClassPath_OneDir")
        }
        // todo
    }
}
