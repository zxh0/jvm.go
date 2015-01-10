package cmdline

import "testing"

func TestParseCommand(t *testing.T) {
    args := []string{"java"}
    _, err := ParseCommand(args)
    if err == nil {
        t.Error("err == nil")
    } else {
        //t.Error(err)
    }
}
