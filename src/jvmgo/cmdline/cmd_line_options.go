package cmdline

import (
    "jvmgo/classpath"
)

type Options struct {
    classpath       *classpath.ClassPath
    verboseClass    bool
}

func (self *Options) Classpath() (*classpath.ClassPath) {
    if self.classpath == nil {
        self.classpath = classpath.ParseClassPath(".")
    }
    return self.classpath
}

func parseOptions(args *CmdLineArgs) (*Options) {
    options := &Options{}

    for !args.empty() && args.first()[0] == '-' {
        optionName := args.removeFirst()
        options.parseClassPathOption(optionName, args)
        // todo
    }

    return options
}

func (self *Options) parseClassPathOption(optionName string, args *CmdLineArgs) bool {
    if optionName == "-classpath" || optionName == "-cp" {
        optionVal := args.removeFirst()
        self.classpath = classpath.ParseClassPath(optionVal)
        return true
    }
    return false
}
