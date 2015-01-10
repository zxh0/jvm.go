package cmdline

import "fmt"

// java [ options ] class [ arguments ]
type Command struct {
    options     []*Option
    class       string
    args        []string
}

// getters
func (self *Command) Class() (string) {
    return self.class
}
func (self *Command) Args() ([]string) {
    return self.args
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

func PrintUsage() {
    fmt.Println("usage: jvmgo [-options] class [args...]")
}
