package cmdline

import (
    "jvmgo/classpath"
)

type Options struct {
    classpath       *classpath.ClassPath
    verboseClass    bool
}

// getters
func (self *Options) Classpath() (*classpath.ClassPath) {
    if self.classpath == nil {
        self.classpath = classpath.ParseClassPath(".")
    }
    return self.classpath
}
func (self *Options) VerboseClass() bool {
    return self.verboseClass
}

func parseOptions(args *CmdLineArgs) (*Options) {
    options := &Options{}

    for !args.isEmpty() && args.first()[0] == '-' {
        optionName := args.removeFirst()
        _ = options.parseClassPathOption(optionName, args) ||
            options.parseVerboseOption(optionName, args)
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

func (self *Options) parseVerboseOption(optionName string, args *CmdLineArgs) bool {
    if optionName == "-verbose" || optionName == "-verbose:class" {
        self.verboseClass = true
        return true
    }
    return false
}
