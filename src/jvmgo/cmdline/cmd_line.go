package cmdline 

//import "fmt"

// java [ options ] class [ arguments ]
type Command struct {
    options []*Option
    class   string
    args    []string
}

type Option struct {
    name    string
    value   string
}

func (self *Command) parseOptions() {

}

func (self *Command) parseClass() {

}

func (self *Command) parseArgs() {

}

func ParseCommand(cmdLineArgs []string) {
    // todo
    cmd := &Command{}
    cmd.options = []*Option{} // len == 0
    cmd.parseOptions()
    cmd.parseClass()
    cmd.parseArgs()
}
