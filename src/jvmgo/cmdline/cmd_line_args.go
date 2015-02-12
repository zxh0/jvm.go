package cmdline

type CmdLineArgs struct {
    args []string
}

func (self *CmdLineArgs) isEmpty() (bool) {
    return len(self.args) == 0
}

func (self *CmdLineArgs) first() (string) {
    return self.args[0]
}

func (self *CmdLineArgs) removeFirst() (first string) {
    first = self.args[0]
    self.args = self.args[1:]
    return
}
