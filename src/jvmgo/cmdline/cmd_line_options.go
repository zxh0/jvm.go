package cmdline

import (
    "strings"
    "jvmgo/classpath"
)

type Options struct {
    classpath *classpath.ClassPath
}

func (self *Options) Classpath() (*classpath.ClassPath) {
    if self.classpath == nil {
        self.classpath = classpath.ParseClassPath(".")
    }
    return self.classpath
}

func parseOptions(args *CmdLineArgs) (*Options) {
    options := &Options{}

    for {
        if args.empty() {
            break
        }
        if !strings.HasPrefix(args.first(), "-") {
            break
        }
        optionName := args.removeFirst()
        if optionName == "-classpath" || optionName == "-cp" {
            optionVal := args.removeFirst()
            options.classpath = classpath.ParseClassPath(optionVal)
        }
        // todo
    }

    return options
}
