package cmdline 

import "fmt"

var options = map[string]bool{"-cp": true, "-classpath": true}

// java [ options ] class [ arguments ]
type Command struct {
    options     []*Option
    class       string
    args        []string
}

type Option struct {
    name    string
    value   string
}

type CmdLineArgs struct {
    args []string
}

func (self *CmdLineArgs) empty() (bool) {
    return len(self.args) > 0
}

func (self *CmdLineArgs) first() (string) {
    return self.args[0]
}

func (self *CmdLineArgs) removeFirst() (first string) {
    first = self.args[0]
    self.args = self.args[1:]
    return
}

func (self *Command) parseOptions(args *CmdLineArgs) {
    if !args.empty() {
        hasVal, isOption := options[args.first()]
        if isOption {
            option := &Option{}
            option.name = args.removeFirst()
            if hasVal {
                option.value = args.removeFirst()
            }

            self.options = append(self.options, option)
            self.parseOptions(args)
        }
    }
}

func (self *Command) parseClass(args *CmdLineArgs) {
    self.class = args.removeFirst()
}

func (self *Command) parseArgs(args *CmdLineArgs) {
    self.args = args.args
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

    cmdLineArgs := &CmdLineArgs{osArgs[1:]}
    cmd = &Command{}
    cmd.options = []*Option{} // len == 0
    cmd.parseOptions(cmdLineArgs)
    cmd.parseClass(cmdLineArgs)
    cmd.parseArgs(cmdLineArgs)
    return
}
