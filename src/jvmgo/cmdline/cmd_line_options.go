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
        if optionName == "-classpath" || optionName == "-cp" {
            optionVal := args.removeFirst()
            options.classpath = classpath.ParseClassPath(optionVal)
        }
        // todo
    }

    return options
}
