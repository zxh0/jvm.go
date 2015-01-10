package cmdline

import "testing"

func TestParseCommandFail0(t *testing.T) {
    args := []string{"java"}
    _, err := ParseCommand(args)
    if err == nil {
        t.Error("err == nil")
    } else {
        //t.Error(err)
    }
}

func TestParseCommandFail1(t *testing.T) {
    args := []string{"java", "-cp"}
    _, err := ParseCommand(args)
    if err == nil {
        t.Error("err == nil")
    } else {
        //t.Error(err)
    }
}

func TestParseCommandFail2(t *testing.T) {
    args := []string{"java", "-cp", "a/b/c"}
    _, err := ParseCommand(args)
    if err == nil {
        t.Error("err == nil")
    } else {
        //t.Error(err)
    }
}

func TestParseCommandOK0(t *testing.T) {
    args := []string{"java", "-cp", "a/b/c", "MyClass"}
    _, err := ParseCommand(args)
    if err != nil {
        t.Error(err)
    } else {
        //t.Error(err)
    }
}

func TestParseCommandOK1(t *testing.T) {
    args := []string{"java", "-cp", "a/b/c", "MyClass", "a", "b", "c"}
    _, err := ParseCommand(args)
    if err != nil {
        t.Error(err)
    } else {
        //t.Error(err)
    }
}
