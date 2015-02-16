package cmdline

import (
	. "jvmgo/test"
	"testing"
)

func TestParseCommandFail0(t *testing.T) {
	args := []string{"java"}
	_, err := ParseCommand(args)
	AssertNotNil(err)
}

func TestParseCommandFail1(t *testing.T) {
	args := []string{"java", "-cp"}
	_, err := ParseCommand(args)
	AssertNotNil(err)
}

func TestParseCommandFail2(t *testing.T) {
	args := []string{"java", "-cp", "a/b/c"}
	_, err := ParseCommand(args)
	AssertNotNil(err)
}

func TestParseCommandOK0(t *testing.T) {
	args := []string{"java", "-cp", "a/b/c", "p.MyClass"}
	cmd, err := ParseCommand(args)
	AssertNil(err)
	AssertNotNil(cmd.options)
	AssertEquals("p/MyClass", cmd.class)
	AssertEquals(0, len(cmd.args))
}

func TestParseCommandOK1(t *testing.T) {
	args := []string{"java", "-cp", "rt.jar", "MyClass", "a", "b", "c"}
	cmd, err := ParseCommand(args)
	AssertNil(err)
	AssertEquals(3, len(cmd.args))
}
