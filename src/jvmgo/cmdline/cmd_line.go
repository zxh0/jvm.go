package cmdline

import (
    "fmt"
    "strings"
)

// java [ options ] class [ arguments ]
type Command struct {
    options *Options
    class   string
    args    []string
}

// getters
func (self *Command) Class() (string) {
    return self.class
}
func (self *Command) Args() ([]string) {
    return self.args
}
func (self *Command) Options() (*Options) {
    return self.options
}

func ParseCommand(osArgs []string) (cmd *Command, err error) {
    defer func() {
        if r := recover(); r != nil {
            var ok bool
            err, ok = r.(error)
            if !ok {
                err = fmt.Errorf("%v", r)
            }
        }
    }()

    args := &CmdLineArgs{osArgs[1:]}
    cmd = &Command{}
    cmd.options = parseOptions(args)
    cmd.class = strings.Replace(args.removeFirst(), ".", "/", -1)
    cmd.args = args.args
    return
}

func PrintUsage() {
    fmt.Println("usage: jvmgo [-options] class [args...]")
}
