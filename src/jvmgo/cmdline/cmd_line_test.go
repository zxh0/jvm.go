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
