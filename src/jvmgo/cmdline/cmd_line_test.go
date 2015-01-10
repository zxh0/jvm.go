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
    cmd, err := ParseCommand(args)
    if err != nil {
        t.Error(err)
    }
    if len(cmd.options) != 1 ||
        cmd.options[0].name != "-classpath" ||
        cmd.options[0].value != "a/b/c" ||
        cmd.class != "MyClass" ||
        len(cmd.args) != 0 {
        t.Error(cmd)
    }
}

func TestParseCommandOK1(t *testing.T) {
    args := []string{"java", "-cp", "a/b/c", "MyClass", "a", "b", "c"}
    cmd, err := ParseCommand(args)
    if err != nil {
        t.Error(err)
    }
    if len(cmd.args) != 3 {
        t.Error(cmd)
    }
}
