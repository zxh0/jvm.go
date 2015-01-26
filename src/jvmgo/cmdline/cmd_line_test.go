package cmdline

import (
    "testing"
    "jvmgo/test"
)

func TestParseCommandFail0(t *testing.T) {
    args := []string{"java"}
    _, err := ParseCommand(args)
    test.AssertNotNil(err)
}

func TestParseCommandFail1(t *testing.T) {
    args := []string{"java", "-cp"}
    _, err := ParseCommand(args)
    test.AssertNotNil(err)
}

func TestParseCommandFail2(t *testing.T) {
    args := []string{"java", "-cp", "a/b/c"}
    _, err := ParseCommand(args)
    test.AssertNotNil(err)
}

func TestParseCommandOK0(t *testing.T) {
    args := []string{"java", "-cp", "a/b/c", "p.MyClass"}
    cmd, err := ParseCommand(args)
    if err != nil {
        t.Error(err)
    }
    if cmd.options == nil || 
            cmd.class != "p/MyClass" || 
            len(cmd.args) != 0 {

        t.Error(cmd)
    }
}

func TestParseCommandOK1(t *testing.T) {
    args := []string{"java", "-cp", "rt.jar", "MyClass", "a", "b", "c"}
    cmd, err := ParseCommand(args)
    if err != nil {
        t.Error(err)
    }
    if len(cmd.args) != 3 {
        t.Error(cmd)
    }
}
