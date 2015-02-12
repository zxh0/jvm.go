package options

import "jvmgo/cmdline"

var VerboseClass bool

func Init(options *cmdline.Options) {
    VerboseClass = options.VerboseClass()
}
